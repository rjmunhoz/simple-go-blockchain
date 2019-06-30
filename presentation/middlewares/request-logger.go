package middlewares

import (
  "fmt"
  "github.com/labstack/echo/v4"
  "time"
)

func RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) (err error) {
    req := c.Request()
    res := c.Response()

    start := time.Now()
    err = next(c); if err != nil {
      c.Error(err)
    }
    stop := time.Now()

    log := fmt.Sprintf("[%s] %s responded %d in %v\n", req.Method, req.URL, res.Status, stop.Sub(start))

    _, err = c.Logger().Output().Write([]byte(log))
    return
  }
}
