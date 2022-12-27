package main

import (
	"context"
	"fmt"
	"github.com/fengzhu0601/goproject/go_tool/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"log"
	"net"
	"server/rpc"
)

func main() {
	logger.InitLogger("../log/server.log", true)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))

	if err != nil {
		log.Fatalf("启动grpc server失败")
		return
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(LogUnaryServerInterceptor(), AuthUnaryServerInterceptor())))
	//grpcServer := grpc.NewServer(grpc.UnaryInterceptor(LogUnaryServerInterceptor()), grpc.UnaryInterceptor(LogUnaryServerInterceptor()))
	//grpcServer := grpc.NewServer()

	rpc.RegisterServerServer(grpcServer, &Server{})

	log.Println("service start")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("启动grpc server失败")
	}
}

func LogUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		logger.Info("method", info.FullMethod, "LogUnaryServerInterceptor")
		resp, err = handler(ctx, req)
		return resp, err
	}
}

func AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		logger.Info("method", info.FullMethod, "AuthUnaryServerInterceptor")
		resp, err = handler(ctx, req)
		return resp, err
	}
}
