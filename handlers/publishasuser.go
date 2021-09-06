package handlers

import (
	"chat/database"
	"chat/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// publishAsUser godoc
// @Summary publishes as user
// @Description Inserts message into database and publishes it to subscribers
// @Tags Messages
// @Accept json
// @Produce  json
// @Param message body models.UserMessageReq true "Message"
// @Security ApiKeyAuth
// @Success 200 {object} models.StatusSuccess
// @Failure 400,500 {object} models.StatusError
// @Failure 401 "Unauthorized"
// @Router /user/message [post]
func publishAsUser(c *gin.Context) {
	var messageReq models.UserMessageReq
	err := c.ShouldBindJSON(&messageReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "provided bad json body",
		})
		return
	}

	if messageReq.Message == "" {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "message is empty",
		})
		return
	}

	userInterface, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusInternalServerError, models.StatusError{
			Error: "bad context user",
		})
		return
	}

	user, ok := userInterface.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.StatusError{
			Error: "bad user interface",
		})
		return
	}

	fmt.Println(user.Id)
	messageID, err := database.InsertMessageGetID(messageReq.Message, user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: err.Error(),
		})
		return
	}
	lmi.SetMessageID(messageID)

	messages := []models.Message{
		{
			MessageID:    messageID,
			Message:      messageReq.Message,
			Username:     user.Username,
			IsRegistered: true,
			SentAt:       time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	messageRes := models.MessageRes{
		LastMessageID: messageID,
		Messages:      messages,
	}

	c.JSON(http.StatusOK, models.StatusSuccess{
		Success: "published as user",
	})
	ps.Publish("global", messageRes)
}
