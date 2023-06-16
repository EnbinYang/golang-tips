package main

import (
	"context"
	"fmt"
	pb "grammar/grpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type CallbackServer struct {
	*pb.UnimplementedCallbackServiceServer
}

func (s *CallbackServer) Callback(ctx context.Context, req *pb.CallbackReq) (*pb.CallbackRsp, error) {
	fmt.Printf("Received callback request for task id: %s\n", req.TaskId)
	return &pb.CallbackRsp{Code: 1}, nil
}

func main() {
	// 创建 gRPC 服务器
	server := grpc.NewServer()

	// 注册回调服务
	pb.RegisterCallbackServiceServer(server, &CallbackServer{})

	// 监听端口并启动服务器
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
