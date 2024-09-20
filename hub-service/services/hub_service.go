package services

import (
	"JINs-software/MOW_SYSTEMS/hub-service/grpc"
	"JINs-software/MOW_SYSTEMS/hub-service/models"
	"context"
	"errors"
)

// HubService는 gRPC 서버에 매치룸 관련 로직을 제공하는 서비스입니다.
type HubService struct {
	grpc.UnimplementedHubServiceServer
}

// CreateMatchRoom은 새로운 매치룸을 생성하는 함수입니다.
func (s *HubService) CreateMatchRoom(ctx context.Context, req *grpc.CreateMatchRoomRequest) (*grpc.CreateMatchRoomResponse, error) {
	// 중복된 이름 확인
	for _, room := range models.MatchRooms {
		if room.RoomName == req.RoomName {
			return nil, errors.New("room name already exists")
		}
	}

	// 새로운 매치룸 생성
	models.LastRoomID++
	room := &models.MatchRoom{
		RoomID:     models.LastRoomID,
		RoomName:   req.RoomName,
		MaxPlayers: int(req.MaxPlayers),
	}
	models.MatchRooms[room.RoomID] = room

	return &grpc.CreateMatchRoomResponse{RoomId: int32(room.RoomID)}, nil
}

// GetMatchRooms는 모든 매치룸 목록을 반환하는 함수입니다.
func (s *HubService) GetMatchRooms(ctx context.Context, req *grpc.GetMatchRoomsRequest) (*grpc.GetMatchRoomsResponse, error) {
	var rooms []*grpc.MatchRoomInfo
	for _, room := range models.MatchRooms {
		rooms = append(rooms, &grpc.MatchRoomInfo{
			RoomId:         int32(room.RoomID),
			RoomName:       room.RoomName,
			MaxPlayers:     int32(room.MaxPlayers),
			CurrentPlayers: int32(room.CurrentPlayers),
		})
	}
	return &grpc.GetMatchRoomsResponse{Rooms: rooms}, nil
}
