package handlers

import (
	"chat/database"
	"chat/models"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// publishAsGuest godoc
// @Summary Publishes message as guest
// @Description Inserts data into database and publishes it to subscribers
// @Tags Messages
// @Accept json
// @Produce  json
// @Param message body models.GuestMessageReq true "Inserts and publishes as guest account"
// @Success 200 {object} models.StatusSuccess
// @Failure 400,500 {object} models.StatusError
// @Router /guest/message [post]
func publishAsGuest(c *gin.Context) {
	var messageReq models.GuestMessageReq
	err := c.ShouldBindJSON(&messageReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "provided bad json body",
		})
		return
	}
	if messageReq.Message == "" || messageReq.Username == "" {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "message or username is empty",
		})
		return
	}
	user, err := database.GetUserByUsername(messageReq.Username)

	//create new user if not exist
	if err == sql.ErrNoRows {
		user.Username = messageReq.Username
		user.Password = ""
		user.Id, err = database.AddUserGetID(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.StatusError{
				Error: err.Error(),
			})
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.StatusError{
			Error: err.Error(),
		})
		return
	}
	if user.Password != "" {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "username is registered",
		})
		return
	}
	id, err := database.InsertMessageGetID(messageReq.Message, user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: err.Error(),
		})
		return
	}
	lmi.SetMessageID(id)

	messages := []models.Message{
		{
			MessageID:    id,
			Message:      messageReq.Message,
			Username:     messageReq.Username,
			IsRegistered: false,
			SentAt:       time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	messageRes := models.MessageRes{
		LastMessageID: id,
		Messages:      messages,
	}

	ps.Publish("global", messageRes)
	c.JSON(http.StatusOK, models.StatusSuccess{
		Success: "published as guest",
	})
}
