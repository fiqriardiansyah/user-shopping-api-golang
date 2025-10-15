package config

import (
	"fmt"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewGrpcServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	return grpcServer
}

func StartGrpcServer(server *grpc.Server) {
	port := os.Getenv("GRPC_USER_CONNECTION_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	logrus.Infof("GRPC server running on port: %s", port)

	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
