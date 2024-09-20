package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config 구조체는 설정 정보를 담습니다.
type Config struct {
	APIPort     string // API Gateway 포트
	AuthService string // AuthService gRPC 서버 주소
	HubService  string // HubService gRPC 서버 주소
}

// LoadConfig 함수는 설정 정보를 로드합니다.
func LoadConfig() Config {
	// .env 파일을 로드합니다.
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found, using default environment variables")
	}

	config := Config{
		APIPort:     getEnv("API_PORT", "3000"),                // 기본적으로 3000 포트를 사용
		AuthService: getEnv("AUTH_SERVICE", "localhost:50052"), // AuthService 주소
		HubService:  getEnv("HUB_SERVICE", "localhost:50051"),  // HubService 주소
	}

	return config
}

// getEnv 함수는 환경 변수를 읽고, 없을 경우 기본값을 반환합니다.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
