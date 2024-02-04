package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sano11o1/snack_logger/logger"
	loggerpb "github.com/sano11o1/snack_logger/logger/grpc"
)

func main() {
	// GRPCサーバーに接続する
	client, conn := logger.InitSnackLoggerConnection()

	e := echo.New()

	// 全てのリクエストでクライアントコネクションを使用できるように設定
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("grpcClient", client)
			return next(c)
		}
	})

	e.Use(logger.SnackLoggerMiddleware)

	e.GET("/hello", HelloHandler)

	e.Logger.Fatal(e.Start(":1324"))

	defer conn.Close()
}

func HelloHandler(c echo.Context) error {
	stream, err := c.Get("stream").(loggerpb.LogService_LogClient)
	if !err {
		return fmt.Errorf("failed to get stream")
	}

	for i := 0; i < 10; i++ {
		if err := logger.Info(fmt.Sprintf("Hello, World! %d", i), stream); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	return c.String(http.StatusOK, "OK!")
}
