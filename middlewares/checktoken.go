package middlewares

import (
	"fmt"
	"net/http"

	"chat/models"
	"chat/redis"

	"github.com/gin-gonic/gin"
)

func CheckToken(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if len(token) < 35 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	username, err := redis.GetStr(token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	id, err := redis.GetInt(username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if username == "" || id == 0 {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user := models.User{
		Id:       id,
		Username: username,
	}
	redis.Setex(username, 30*60, fmt.Sprintf("%d", id))
	redis.Setex(token, 30*60, username)
	c.Set("user", user)
}
