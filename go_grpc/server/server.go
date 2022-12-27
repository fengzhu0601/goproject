package main

import (
	"context"
	"fmt"
	"github.com/fengzhu0601/goproject/go_tool/logger"
	"server/rpc"
)

type Server struct {
}

func (s *Server) Hello(ctx context.Context, request *rpc.Empty) (*rpc.HelloResponse, error) {
	logger.Info("call hello")
	resp := &rpc.HelloResponse{
		Hello: "hello client",
	}
	return resp, nil
}

func (s *Server) Register(ctx context.Context, request *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	logger.Info("call register")
	resp := &rpc.RegisterResponse{
		Uid: fmt.Sprintf("%s.%s", request.GetName(), request.GetPassword()),
	}
	return resp, nil
}
