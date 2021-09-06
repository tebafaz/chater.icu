package handlers

import (
	"chat/models"
	"chat/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

// logout godoc
// @Summary Logs out
// @Description Logs out user by deleting his token in server
// @Tags Token
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.StatusSuccess
// @Failure 400 {object} models.StatusError
// @Failure 401 "Unauthorized"
// @Router /user/logout [post]
func logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "token is not provided",
		})
		return
	}
	err := redis.Del(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "you are not logged in",
		})
		return
	}
	c.JSON(http.StatusOK, models.StatusSuccess{
		Success: "logged out",
	})
}
