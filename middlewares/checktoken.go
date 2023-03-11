package middlewares

import (
	"fmt"
	"net/http"

	"github.com/tebafaz/chater.icu/models"
	"github.com/tebafaz/chater.icu/redis"

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
	err = redis.Setex(username, 30*60, fmt.Sprintf("%d", id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = redis.Setex(token, 30*60, username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Set("user", user)
}
