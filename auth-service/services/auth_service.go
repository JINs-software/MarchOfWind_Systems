package services

import (
	"JINs-software/MOW_SYSTEMS/auth-service/models"
	grpc "JINs-software/MOW_SYSTEMS/proto"
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"database/sql"

	"crypto/rand"
	"encoding/hex"

	"github.com/redis/go-redis/v9"
)

// 인터페이스 도입, 서비스 코드에서는 실제 sqlDB나 redis 객체에 접근
// 테스트 코드에서는 Mock 객체 사용. 따라서 Mock 객체가 일반적인 DB나 Redis와 같은 인터페이스를 갖도록 함

// [인터페이스]
//	- 인터페이스 선언
//		type InterfaceName interface {
//			Fly();					// 인테페이스 메서드 0
//			Walk(distance int) int	// 인터페이스 메서드 1
//		}
//
//	- type 키워드: 인터페이스도 구조체처럼 타입 중 하나, 변수 선언 가능, 변수의 값으로 사용할 수 있음
// 	- 메서드 집합
//		(1) 메서드는 반드시 메서드 명이 존재해야 함
//		(2) 매개변수와 반환이 다르더라도 이름이 같은 메서드를 가질 수 없음
//		(3) 인터페이스에는 메서드 구현을 포함하지 않음
//
//// 1. Stringer 인터페이스 선언
//// Stringer 인터페이스는 매개변수 없는 string 타입을 반환하는 String() 메서드를 포함
//// -> string 타입을 반환하는 String() 메서드를 포함한 모든 타입은 Stringer 인터페이스로 사용될 수 있음!!
//type Stringer interface {
//	String() string
//}
//
//type Student struct {
//	Name string
//	Age  int
//}
//
//// 2. Student의 String() 메서드
//// Student 타입은 String() 메서드를 포함, 따라서 Student 타입은 Stringer 인터페이스로 사용 가능!
//func (s Student) String() string {
//	return s.Name
//}
//
//func sample() {
//	student := Student("철수", 123)
//	var stringer Stringer
//	stringer = student // stringer 값으로 student 대입
//	// Student 타입은 String() 메서드를 포함하고 있기에 대입 가능하다.
//	print(stringer.String())
//}

// DB 인터페이스 정의
type DB interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Close() error
}

// Redis 인터페이스 정의
type Redis interface {
	Ping(ctx context.Context) *redis.StatusCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

var (
	syncOnce sync.Once
	//DbConn    *sql.DB
	//RedisConn *redis.Client
	//Ctx       = context.Background()
	// => go 에서 소문자로 시작하는 변수는 외부 패키지에서 접근할 수 없음

	//DbConn    *sql.DB
	//RedisConn *redis.Client
	// => 인터페이스 활용

	// Init 함수에서 sql.Open(..) 함수를 통해 db 연결 객체를 할당받아 대입
	// sqlDB는 이미 DB 인터페이스 메서드 집합을 정의한 상태이다.
	DbConn DB

	// DB 인터페이스와 마찬가지
	RedisConn Redis
	Ctx       = context.Background()
)

type AuthService struct {
	grpc.UnimplementedAuthServiceServer
}

func generateToken(username string) string { // <- 경고 발생
	// 16바이트의 랜덤 데이터를 생성하여 고유한 토큰으로 사용
	tokenBytes := make([]byte, 16)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "" // 에러가 발생하면 빈 문자열 반환
	}
	// username을 토큰의 앞에 추가
	return username + "-" + hex.EncodeToString(tokenBytes)
}

func Init() {
	syncOnce.Do(func() {
		var err error

		// Connect To MySQL Server
		DbConn, err = sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/accountDB")
		//DbConn, err = sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/accountDB")
		if err != nil {
			log.Fatalf("Failed to Connect to AccountDB: %v", err)
		}

		// Connect To Redis (Token MemDB)
		RedisConn = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		_, err = RedisConn.Ping(Ctx).Result()
		if err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
		}
	})
}

// 참고, grpc 구조체임
//
//	type CreateAccountRequest struct {
//		state         protoimpl.MessageState
//		sizeCache     protoimpl.SizeCache
//		unknownFields protoimpl.UnknownFields
//
//		Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
//		Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
//	}
//
//	type CreateAccountResponse struct {
//		state         protoimpl.MessageState
//		sizeCache     protoimpl.SizeCache
//		unknownFields protoimpl.UnknownFields
//
//		Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
//	}
func (s *AuthService) CreateAccount(Ctx context.Context, req *grpc.CreateAccountRequest) (*grpc.CreateAccountResponse, error) {
	// 1) DB에 중복된 이름이 있는지 확인
	query := "SELECT username FROM users WHERE username = ?"
	row := DbConn.QueryRowContext(Ctx, query, req.Username)
	var existingUsername string
	err := row.Scan(&existingUsername)

	if err != nil {
		if err == sql.ErrNoRows {
			// 중복된 이름이 없음
		} else {
			// 다른 DB 관련 에러
			return &grpc.CreateAccountResponse{
				Status: 500, // 내부 서버 오류
			}, errors.New("database error: " + err.Error())
		}
	} else {
		// 중복된 이름 존재 시 실패 상태 반환
		return &grpc.CreateAccountResponse{
			Status: 409, // 충돌 상태
		}, errors.New("username already exists")
	}

	// 2) 비밀번호 해시화 (필요시)
	//hashedPassword, err := HashPassword(req.Password)
	hashedPassword := req.Password

	if err != nil {
		return &grpc.CreateAccountResponse{
			Status: 500, // 내부 서버 오류
		}, errors.New("failed to hash password: " + err.Error())
	}

	// 3) 계정 생성 쿼리
	insertQuery := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err = DbConn.ExecContext(Ctx, insertQuery, req.Username, hashedPassword)
	if err != nil {
		return &grpc.CreateAccountResponse{
			Status: 500, // 내부 서버 오류
		}, errors.New("failed to create account: " + err.Error())
	}

	// 4) 성공 상태 반환
	return &grpc.CreateAccountResponse{
		Status: 200, // 성공 상태
	}, nil
}

func (s *AuthService) Login(Ctx context.Context, req *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	var err error
	// 1) db를 통해 user 정보 존재 확인
	// 2) 만약 존재하지 않거나, 비밀번호가 틀리다면 실패를 반환 (Token을 null로 전달)
	// 3) 만약 존재한다면 토큰을 발급하고, models.Users 맵에 {이름, 패스워드, 발급한 토큰} 저장

	query := "SELECT password FROM users WHERE username = ?"
	row := DbConn.QueryRowContext(Ctx, query, req.Username)
	var storedPassword string
	err = row.Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			// 유저가 존재하지 않음
		}

		// 다른 DB 관련 에러
		return &grpc.LoginResponse{
			Token: "",
		}, errors.New("non-existent user information") // <- 경고
	}

	// 2. password 확인
	if storedPassword != req.Password {
		return &grpc.LoginResponse{
			Token: "",
		}, errors.New("invalid credentials")
	}

	// 3. 토큰 발급
	token := generateToken(req.Username)

	// 4. Redis에 key: req.Username, value: 발급한 토큰(token)을 삽입해줘
	err = RedisConn.Set(Ctx, req.Username, token, 0).Err()
	if err != nil {
		log.Fatalf("failed to set token in Redis: %v", err)
	}

	// 5. 캐싱
	models.Users[req.Username] = models.User{
		Username: req.Username,
		Password: storedPassword,
		Token:    token,
	}

	return &grpc.LoginResponse{
		Token: token,
	}, nil
}
