package services

import (
	"JINs-software/MOW_SYSTEMS/auth-service/grpc"
	"JINs-software/MOW_SYSTEMS/auth-service/models"
	"context"
	"errors"
	"log"
	"sync"

	"database/sql"

	"crypto/rand"
	"encoding/hex"

	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	grpc.UnimplementedAuthServiceServer
}

var (
	syncOnce  sync.Once
	dbConn    *sql.DB
	redisConn *redis.Client
	ctx       = context.Background()
)

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
		dbConn, err = sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/accountDB")
		if err != nil {
			log.Fatalf("Failed to Connect to AccountDB: %v", err)
		}

		// Connect To Redis (Token MemDB)
		redisConn = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		_, err = redisConn.Ping(ctx).Result()
		if err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
		}
	})
}

func (s *AuthService) Login(ctx context.Context, req *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	var err error
	// 1) db를 통해 user 정보 존재 확인
	// 2) 만약 존재하지 않거나, 비밀번호가 틀리다면 실패를 반환 (Token을 null로 전달)
	// 3) 만약 존재한다면 토큰을 발급하고, models.Users 맵에 {이름, 패스워드, 발급한 토큰} 저장

	query := "SELECT password FROM users WHERE username = ?"
	row := dbConn.QueryRowContext(ctx, query, req.Username)
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
	err = redisConn.Set(ctx, req.Username, token, 0).Err()
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
