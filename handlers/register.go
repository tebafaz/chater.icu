package handlers

import (
	"fmt"
	"net/http"

	"github.com/tebafaz/chater.icu/database"
	"github.com/tebafaz/chater.icu/helpers"
	"github.com/tebafaz/chater.icu/models"
	"github.com/tebafaz/chater.icu/redis"

	"github.com/gin-gonic/gin"
)

// registerAndLogin godoc
// @Summary Register and get token
// @Description Registers new user and returns token. Token lasts 30 minutes and updates it at each performed action with token
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
	err = redis.Setex(tokenResponse.Token, 30*60, user.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = redis.Setex(user.Username, 30*60, fmt.Sprintf("%d", id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tokenResponse)
}
