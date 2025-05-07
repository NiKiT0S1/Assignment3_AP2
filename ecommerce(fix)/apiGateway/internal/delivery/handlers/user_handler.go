package handlers

import (
	grpcDelivery "apiGateway/internal/grpc"
	"apiGateway/internal/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, userClient *grpcDelivery.UserClient) {
	r.POST("/register", func(c *gin.Context) {
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}
		user, err := userClient.Register(body.Username, body.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "registration failed: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": user.Id, "username": user.Username})
	})

	r.POST("/login", func(c *gin.Context) {
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		// Создаем запрос на авторизацию
		req := &proto.AuthRequest{
			Username: body.Username,
			Password: body.Password,
		}

		user, err := userClient.Authenticate(c, req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user_id": user.Id, "token": "Bearer " + strconv.Itoa(int(user.Id))})
	})

	r.GET("/profile/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}

		user, err := userClient.GetProfile(int32(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": user.Id, "username": user.Username})
	})
}
