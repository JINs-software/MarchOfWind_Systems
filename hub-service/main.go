package main

import (
	hubRPC "JINs-software/MOW_SYSTEMS/hub-service/grpc"
	"JINs-software/MOW_SYSTEMS/hub-service/services"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// TCP 포트 50051에서 리스닝
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	// gRPC 서버 생성
	server := grpc.NewServer()
	hubRPC.RegisterHubServiceServer(server, &services.HubService{})
	reflection.Register(server) // gRPC reflection을 사용하여 클라이언트에서 서비스 탐색 가능

	println("HubService is running on port 50051...")
	server.Serve(lis)
}
