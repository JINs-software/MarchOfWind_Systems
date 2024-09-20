package services

import (
	"context"
	"fmt"
	"log"
	"sync"

	hubRpc "JINs-software/MOW_SYSTEMS/api-gateway/grpc" // hub_service_grpc 패키지에 별칭 'hubRpc' 사용

	"google.golang.org/grpc" // gRPC 패키지 임포트
)

var (
	hubClient hubRpc.HubServiceClient // gRPC 클라이언트를 전역 변수로 관리
	hubConn   *grpc.ClientConn        // gRPC 연결 객체 전역 관리
	hubOnce   sync.Once               // 한 번만 연결을 생성하기 위한 sync.Once
)

// gRPC 연결 초기화 함수
func initHubServiceClient() {
	hubOnce.Do(func() {
		var err error
		// gRPC 서버에 연결 (HubService 서버가 50051 포트에서 실행된다고 가정)
		hubConn, err = grpc.Dial("localhost:50051", grpc.WithInsecure()) // 실제 gRPC 서버 주소로 변경 필요
		if err != nil {
			log.Fatalf("Failed to connect to HubService: %v", err)
		}
		hubClient = hubRpc.NewHubServiceClient(hubConn) // 'hubRpc' 별칭을 사용하여 클라이언트 생성
	})
}

// GetMatchRooms 함수는 gRPC를 사용하여 매치룸 목록을 가져옵니다.
func GetMatchRooms() ([]map[string]interface{}, error) {
	// gRPC 클라이언트 초기화 (한 번만 실행)
	initHubServiceClient()

	// gRPC 요청을 통해 매치룸 목록을 가져옴
	resp, err := hubClient.GetMatchRooms(context.Background(), &hubRpc.GetMatchRoomsRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch matchrooms: %v", err)
	}

	// 매치룸 목록을 변환하여 반환
	var rooms []map[string]interface{}
	for _, room := range resp.Rooms {
		rooms = append(rooms, map[string]interface{}{
			"roomId":         room.RoomId,
			"roomName":       room.RoomName,
			"maxPlayers":     room.MaxPlayers,
			"currentPlayers": room.CurrentPlayers,
		})
	}
	return rooms, nil
}

// CreateMatchRoom 함수는 gRPC를 사용하여 새로운 매치룸을 생성합니다.
func CreateMatchRoom(roomName string, maxPlayers int) error {
	// gRPC 클라이언트 초기화 (한 번만 실행)
	initHubServiceClient()

	// gRPC 요청을 통해 매치룸 생성
	_, err := hubClient.CreateMatchRoom(context.Background(), &hubRpc.CreateMatchRoomRequest{
		RoomName:   roomName,
		MaxPlayers: int32(maxPlayers),
	})
	if err != nil {
		return fmt.Errorf("failed to create matchroom: %v", err)
	}
	return nil
}
