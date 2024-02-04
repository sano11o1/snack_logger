package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"

	loggerpb "github.com/sano11o1/snack_logger/logger/grpc"
)

type Server struct {
	loggerpb.UnimplementedLogServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	loggerpb.RegisterLogServiceServer(s, NewServer())

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()

}

func (s Server) Log(stream loggerpb.LogService_LogServer) error {
	for {
		req, err := stream.Recv()
		fmt.Println("=============Request============", req.GetMessage())
		if errors.Is(err, io.EOF) {
			fmt.Println("=============EOF============")
			time.Sleep(10 * time.Second)
			return stream.SendAndClose(&loggerpb.LogResponse{
				Message: "End of Universe",
			})
		}
		if err != nil {
			return err
		}
	}
}
