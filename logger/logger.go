package logger

import (
	"fmt"
	"log"

	client "github.com/sano11o1/snack_logger/logger/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitSnackLogger() (client.LogServiceClient, *grpc.ClientConn) {
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

	return client.NewLogServiceClient(conn), conn
}
