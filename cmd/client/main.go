package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sano11o1/snack_logger/logger"
)

func main() {
	logger.InitSnackLogger()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
