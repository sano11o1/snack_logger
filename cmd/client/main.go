package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	loggerpb "github.com/sano11o1/snack_logger/logger/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// GRPCサーバーに接続する
	client, conn := InitSnackLoggerConnection()

	e := echo.New()

	// 全てのリクエストでクライアントコネクションを使用できるように設定
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("grpcClient", client)
			return next(c)
		}
	})

	e.Use(SnackLoggerMiddleware)

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
		if err := Info(fmt.Sprintf("Hello, World! %d", i), stream); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	return c.String(http.StatusOK, "OK!")
}

func SnackLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		grpcClient, ok := c.Get("grpcClient").(loggerpb.LogServiceClient)
		if !ok {
			return fmt.Errorf("failed to get grpc client")
		}

		stream, err := grpcClient.Log(context.Background())
		if err != nil {
			return err
		}
		c.Set("stream", stream)

		if err := next(c); err != nil {
			return err
		}

		res, err := stream.CloseAndRecv()
		if err != nil {
			return err
		}
		fmt.Println(res.GetMessage())
		return nil
	}
}

func InitSnackLoggerConnection() (loggerpb.LogServiceClient, *grpc.ClientConn) {
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
		return nil, nil
	}

	fmt.Println("start gRPC Client.")

	return loggerpb.NewLogServiceClient(conn), conn
}

func Info(msg string, stream loggerpb.LogService_LogClient) error {
	fmt.Println(msg)

	if err := stream.Send(&loggerpb.LogRequest{
		Message: msg,
	}); err != nil {
		return err
	}

	return nil
}
