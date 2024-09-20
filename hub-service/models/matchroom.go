package models

// MatchRoom 구조체는 매치룸 정보를 담고 있습니다.
type MatchRoom struct {
	RoomID         int    // 매치룸 고유 ID
	RoomName       string // 매치룸 이름
	MaxPlayers     int    // 최대 플레이어 수
	CurrentPlayers int    // 현재 플레이어 수
}

// 매치룸을 관리하는 전역 변수
var MatchRooms = make(map[int]*MatchRoom) // RoomID를 키로 매치룸 정보를 저장
var LastRoomID int                        // 마지막으로 생성된 매치룸의 ID
