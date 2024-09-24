package test

import (
	"JINs-software/MOW_SYSTEMS/auth-service/services"
	grpc "JINs-software/MOW_SYSTEMS/proto"
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// stretchr/testify/assert 패키지 API
// - Equal(): expected와 actual 두 값을 비교, 다를 경우 실패하고 메시지를 출력
// - Greater(): e1이 e2보다 크지 않으면 실패하고 메시지를 출력
// - Len(): object의 항목 개수가 length가 아니면 테스트를 실패하고 메시지를 출력
// - NotNil(): object가 nil이면 테스트를 실패하고 메시지를 출력
// - NotEqualf(): expected와 actual이 같으면 테스트를 실패하고 메시지를 출력

//  stretchr/testify/mock 패키지
//	:	모듈의 행동을 가장하는 mockup 객체를 제공
//		ex) 온라인 기능을 테스트할 때 하위 영역인 네트워크 기능까지 모두 테스트하기 힘들 때,
//			네트워크 객체를 가장하는 목업 객체를 만들 때 유용
// 	stretchr/testify/suite 패키지
//	: 	테스트 준비 작업이나 테스트 종료 후 뒤처리 작업을 쉽게 할 수 있도록 함
//		ex) 테스트에 특정 파일이 있어야 한다면 테스트 시작 전 임시 파일을 생성하고,
//			테스트 종료 후 생성한 임시 파일을 삭제해주는 작업을 만들 떄 유용

// Mock 객체를 활용하여 DB 연결 및 Redis 연결 없이 로직 검증
// MockDB는 DB 연결을 흉내내는 목(Mock) 객체입니다.
// Mock for DB connection

type MockDB struct {
	mock.Mock
}

type MockRow struct {
	//data map[string]interface{} // 올바른 인터페이스 타입
	value string
}

// Scan 메서드 구현
func (m *MockRow) Scan(dest ...interface{}) error {
	if len(dest) == 0 {
		return errors.New("no destination provided")
	}
	// Mock 데이터를 dest에 할당
	*dest[0].(*string) = m.value
	return nil
}

func (m *MockDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	//mockArgs := m.Called(ctx, query, args)
	//return mockArgs.Get(0).(*sql.Row) // MockRow 반환

	// MockRow를 생성하여 값을 설정
	//return &MockRow{value: "mocked_username"} // <- 컴파일 에러, cannot use &MockRow{…} (value of type *MockRow) as *sql.Row value in return statement
}

//type MockRow struct {
//	//data map[string]interface{} // 올바른 인터페이스 타입
//	value string
//}
//
//// Scan 메서드는 쿼리 결과를 주어진 변수에 저장하는 역할을 합니다.
//func (r *MockRow) Scan(dest ...interface{}) error {
//	// 예시로 "username"과 "password"를 반환한다고 가정
//	dest[0] = "testuser"    // 첫 번째 변수에 "testuser" 저장
//	dest[1] = "password123" // 두 번째 변수에 "password123" 저장
//	return nil
//}

func (m *MockDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return nil, nil // Dummy implementation
}
func (m *MockDB) Close() error {
	return nil // Dummy implementation
}

// Mock for Redis connection
type MockRedis struct {
	mock.Mock
}

func (m *MockRedis) Ping(ctx context.Context) *redis.StatusCmd {
	mockArgs := m.Called(ctx)
	return mockArgs.Get(0).(*redis.StatusCmd)
}

func (m *MockRedis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	mockArgs := m.Called(ctx, key, value, expiration)
	return mockArgs.Get(0).(*redis.StatusCmd)
}

func TestInit(t *testing.T) {
	// auth_service의 Init 함수 호출
	// -> DB, Redis 연결 테스트
	services.Init()

	// DB 연결 정상 여부 확인
	assert.NotNil(t, services.DbConn, "account db connection should be initialized")

	// Redis 연결 정상 여부 확인
	assert.NotNil(t, services.RedisConn, "redis connection should be initialized")

	// redis에 ping 명령을 보내 연결이 유지되는지 확인
	pong, err := services.RedisConn.Ping(services.Ctx).Result()
	assert.Nil(t, err, "Ping to Redis should not return an error")
	assert.Equal(t, "PONG", pong, "Redis should respond with PONG")
}

func TestLogin(t *testing.T) {
	// Mock 객체 생성
	mockDB := new(MockDB)
	mockRedis := new(MockRedis)
	// context 생성
	ctx := context.Background()

	// Mock's behavior 설정
	// Mock DB behavior 설정
	// mockDB 객체에서 QueryRowContext 호출 시 username이 "testuser"일 경우 "password123"을 반환하도록 설정
	mockDB.On("QueryRowContext", ctx, "SELECT password FROM users WHERE username = ?", mock.Anything).Return("password123", nil)

	// 테스트할 서비스에 mock 객체 주입(injection)
	services.DbConn = mockDB       // cannot use mockDB (variable of type *MockDB) as *sql.DB value in assignment
	services.RedisConn = mockRedis // cannot use mockDB (variable of type *MockDB) as *sql.DB value in assignment
	services.Ctx = ctx

	// 테스트할 gRPC 요청 생성 (사용자 이름과 비밀번호가 포함된 요청)
	req := &grpc.LoginRequest{Username: "jinhur", Password: "password123"}

	// services.Login 함수 호출 (mock 객체가 사용됨)

	authService := &services.AuthService{}
	resp, err := authService.Login(ctx, req) // <- 컴파일 에러, undefined: services.Login

	// 결과 검증 (assert 라이브러리 사용)
	// 1. 에러가 없어야 함
	assert.Nil(t, err, "error should be nil")

	// 2. 생성된 토큰이 "testuser-"로 시작하는지 확인 (이름 기반 토큰 생성 확인)
	assert.Equal(t, "testuser-", resp.Token[:9], "Token should be generated with username prefix")
}
