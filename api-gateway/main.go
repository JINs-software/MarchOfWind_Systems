package main

import (
	"JINs-software/MOW_SYSTEMS/api-gateway/config"
	"JINs-software/MOW_SYSTEMS/api-gateway/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 설정 로드
	cfg := config.LoadConfig()

	r := gin.Default()

	// 매치룸 관련 핸들러 등록
	r.GET("/matchrooms", handlers.GetMatchRoomsHandler)
	r.POST("/matchroom/create", handlers.CreateMatchRoomHandler)

	// Auth 관련 핸들러 등록
	r.POST("/login", handlers.LoginHandler)

	//r.Run(":3000") // API Gateway는 3000 포트에서 실행
	// API Gateway 서버 실행
	r.Run(":" + cfg.APIPort) // 설정 파일에서 포트를 로드
}
