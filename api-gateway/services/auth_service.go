package services

import (
	"context"
	"fmt"
	"log"
	"sync"

	authRpc "JINs-software/MOW_SYSTEMS/api-gateway/grpc" // auth_service_grpc 패키지 임포트

	"google.golang.org/grpc" // gRPC 패키지 임포트
)

var (
	authClient authRpc.AuthServiceClient // gRPC 클라이언트를 전역 변수로 관리
	authConn   *grpc.ClientConn          // gRPC 연결 객체 전역 관리
	authOnce   sync.Once                 // 한 번만 연결을 생성하기 위한 sync.Once
)

// gRPC 연결 초기화 함수
func initAuthServiceClient() {
	authOnce.Do(func() {
		var err error
		// gRPC 서버에 연결 (AuthService 서버가 50052 포트에서 실행된다고 가정)
		authConn, err = grpc.Dial("localhost:50052", grpc.WithInsecure()) // 실제 gRPC 서버 주소로 변경 필요
		if err != nil {
			log.Fatalf("Failed to connect to AuthService: %v", err)
		}
		authClient = authRpc.NewAuthServiceClient(authConn)
	})
}

// 로그인 함수
func Login(username, password string) (string, error) {
	// gRPC 클라이언트 초기화 (한 번만 실행)
	initAuthServiceClient()

	// gRPC 로그인 요청
	resp, err := authClient.Login(context.Background(), &authRpc.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return "", fmt.Errorf("failed to login: %v", err)
	}

	return resp.Token, nil
}
