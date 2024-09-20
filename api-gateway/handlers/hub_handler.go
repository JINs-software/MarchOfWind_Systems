package handlers

import (
	"JINs-software/MOW_SYSTEMS/api-gateway/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMatchRoomsHandler(c *gin.Context) {
	rooms, err := services.GetMatchRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rooms": rooms})
}

func CreateMatchRoomHandler(c *gin.Context) {
	var req struct {
		RoomName   string `json:"roomName" binding:"required"`
		MaxPlayers int    `json:"maxPlayers" binding:"required,min=2,max=4"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateMatchRoom(req.RoomName, req.MaxPlayers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Matchroom created successfully"})
}
