package handlers

import (
	"chat/database"
	"chat/helpers"
	"chat/models"
	"chat/redis"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// registerAndLogin godoc
// @Summary Register and get token
// @Description Registers new user and returns token
// @Tags Token
// @Accept json
// @Produce  json
// @Param token body models.TokenReq true "Token"
// @Success 200 {object} models.TokenRes
// @Failure 400 {object} models.StatusError
// @Router /register [post]
func registerAndLogin(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "provided bad json data",
		})
		return
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "username or password is empty",
		})
		return
	}
	user.Password, err = helpers.CreateHash(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "could not create hash for password",
		})
		return
	}
	id, err := database.AddUserGetID(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: err.Error(),
		})
		return
	}

	tokenResponse := models.TokenRes{
		Token: helpers.CreateToken(),
	}
	redis.Setex(tokenResponse.Token, 30*60, user.Username)
	redis.Setex(user.Username, 30*60, fmt.Sprintf("%d", id))
	c.JSON(http.StatusOK, tokenResponse)
}
