package routes

import (
  "chain/errors"
  "chain/services"
  "fmt"
  "github.com/labstack/echo/v4"
  "net/http"
  "strconv"
)

func GetBlock(service services.BlockService) echo.HandlerFunc {
  return func(c echo.Context) error {
    sequence, err := strconv.Atoi(c.Param("sequence"))

    if err != nil {
      return c.JSON(errors.NewHttpError(http.StatusUnprocessableEntity, "sequence must be an int", "unprocessable_entity"))
    }

    block := service.Find(sequence)

    if block.Data == "" {
      status, httpError := errors.NewNotFoundError(errors.ErrorDetails{Message: fmt.Sprintf("block `%d` was not found", sequence)})
      return c.JSON(status, httpError)
    }

    return c.JSON(http.StatusOK, block)
  }
}
