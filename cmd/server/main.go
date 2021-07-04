package main

import (
	"grpc-simple/internal/service"
	"grpc-simple/internal/transport/rpc"

	"github.com/sirupsen/logrus"
)

func main() {
	srv := service.NewService()
	h := rpc.NewHandler(srv)
	log := logrus.New()
	rpc.StartRPCServer(h, "7000", log)
}
