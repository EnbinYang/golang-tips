package main

import (
	"context"
	pb "grammar/grpc"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// 建立与 gRPC 服务器的连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewCallbackServiceClient(conn)

	// 创建回调请求
	req := &pb.CallbackReq{
		TaskId: "123456",
		Data:   0,
	}

	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 发起回调
	rsp, err := client.Callback(ctx, req)
	if err != nil {
		log.Fatalf("Failed to send callback request: %v", err)
	}

	log.Printf("Callback response: code=%d", rsp.Code)
}
