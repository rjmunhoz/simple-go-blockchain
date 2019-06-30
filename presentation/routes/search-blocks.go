package routes

import (
  "chain/services"
  "fmt"
  "github.com/labstack/echo/v4"
  "net/http"
  "strconv"
)

func SearchBlocks (service services.BlockService) echo.HandlerFunc {
  return func(c echo.Context) error {
    page, _ := strconv.Atoi(c.QueryParam("page")); if page > 0 {
      page = page -1
    }

    size, err := strconv.Atoi(c.QueryParam("size")); if err != nil {
      size = 10
    }

    result := service.Search(page, size)

    status := http.StatusOK

    if result.Count < result.Total {
      status = http.StatusPartialContent
      contentRange := fmt.Sprintf("%d-%d/%d", result.From, result.From + result.Count, result.Total)
      c.Response().Header().Add("content-range", contentRange)
    }

    return c.JSON(status, result.Results)
  }
}
