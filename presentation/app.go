package presentation

import (
  "chain/config"
  "chain/data/repositories"
  "chain/domain"
  "chain/presentation/middlewares"
  "chain/presentation/routes"
  "chain/services"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

func AppFactory(_ config.AppConfig, blockChain []domain.MinedBlock) *echo.Echo {
  e := echo.New()

  e.Use(middleware.Recover())
  e.Use(middlewares.RequestLogger)

  blockRepository := repositories.BlockRepository{BlockChain: blockChain}
  blockService := services.BlockService{Repository: blockRepository}

  e.GET("/ping", routes.Ping)
  e.GET("/blocks", routes.SearchBlocks(blockService))
  e.GET("/blocks/:sequence", routes.GetBlock(blockService))

  return e
}
