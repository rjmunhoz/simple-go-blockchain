package presentation

import (
  "chain/config"
  "chain/domain"
)

func StartServer(config config.AppConfig, blockChain []domain.MinedBlock) {
  app := AppFactory(config, blockChain)
  app.Logger.Fatal(app.Start(config.Port))
}
