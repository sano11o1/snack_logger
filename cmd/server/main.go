package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"

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
		if errors.Is(err, io.EOF) {
			fmt.Println("request:", "EOF")
			return stream.SendAndClose(&loggerpb.LogResponse{
				Message: "End of Stream",
			})
		}
		fmt.Println("request:", req.GetMessage())
		if err != nil {
			return err
		}
	}
}
