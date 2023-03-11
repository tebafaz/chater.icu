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

// login godoc
// @Summary Login
// @Description Logs in by giving token. Token lasts 30 minutes and updates it at each performed action with token
// @Tags Token
// @Accept json
// @Produce  json
// @Param tokenReq body models.TokenReq true "Username and password"
// @Success 200 {object} models.TokenRes
// @Failure 400 {object} models.StatusError
// @Router /login [post]
func login(c *gin.Context) {
	var tokenReq models.TokenReq
	err := c.ShouldBindJSON(&tokenReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "provided bad json body",
		})
		return
	}
	if tokenReq.Username == "" || tokenReq.Password == "" {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "Username or password is empty",
		})
		return
	}
	user, err := database.GetUserByUsername(tokenReq.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: err.Error(),
		})
		return
	}
	passwdIsEqual := helpers.ComparePasswords(tokenReq.Password, user.Password)
	if !passwdIsEqual {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "password doesn't match",
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
	err = redis.Setex(user.Username, 30*60, fmt.Sprintf("%d", user.Id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tokenResponse)
}
