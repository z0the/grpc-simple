package rpc

import (
	"context"
	"grpc-simple/internal/service"
	"grpc-simple/pkg/api"
	"net"

	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
)

func StartRPCServer(hdl *Handler, port string, log *logrus.Logger) {
	address := ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		logrus.Fatal(err)
	}

	log.Infof("RPC Server is running on port %s...", port)
	err = setupRouter(hdl).Serve(listener)
	if err != nil {
		log.WithError(err).Fatal("failed run server!")
	}
}

func setupRouter(hdl *Handler) *grpc.Server {
	s := grpc.NewServer()
	api.RegisterMathServer(s, hdl)
	return s
}

type Handler struct {
	api.UnimplementedMathServer
	srv *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (s *Handler) Sum(ctx context.Context, req *api.SumReq) (*api.SumResp, error) {
	res := s.srv.Sum(ctx, req.X, req.Y)
	resp := &api.SumResp{Result: res}
	return resp, nil
}

func (s *Handler) Sub(ctx context.Context, req *api.SubReq) (*api.SubResp, error) {
	res := s.srv.Sub(ctx, req.X, req.Y)
	resp := &api.SubResp{Result: res}
	return resp, nil
}
