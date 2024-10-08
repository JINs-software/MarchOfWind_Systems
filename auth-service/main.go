package main

import (
	"JINs-software/MOW_SYSTEMS/auth-service/services"
	authRPC "JINs-software/MOW_SYSTEMS/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	authRPC.RegisterAuthServiceServer(server, &services.AuthService{})
	reflection.Register(server)

	println("AuthService is running on port 50052...")
	server.Serve(lis)
}
