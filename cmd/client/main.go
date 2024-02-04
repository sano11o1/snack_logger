package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sano11o1/snack_logger/logger"
	loggerpb "github.com/sano11o1/snack_logger/logger/grpc"
)

func main() {
	client, conn := logger.InitSnackLogger()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		if err := stream(client); err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Error")
		}
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1324"))

	defer conn.Close()
}

func stream(c loggerpb.LogServiceClient) error {
	fmt.Println("echoサーバーのリクエストでstreamを開始します。")
	client, err := c.Route(context.Background())
	if err != nil {
		return err
	}

	if err := client.Send(&loggerpb.LogRequest{
		Message: "Hello, World!",
	}); err != nil {
		return err
	}

	res, err := client.CloseAndRecv()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetMessage())
	}
	fmt.Println("echoサーバーのリクエストでstreamを終了します。")
	return nil
}
