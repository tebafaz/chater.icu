package handlers

import (
	"chat/database"
	"chat/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getIfExist(c *gin.Context) {
	var messageID models.ID
	err := c.ShouldBind(&messageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "provided bad json body",
		})
		c.Abort()
		return
	}

	lastMessageID := lmi.GetMessageID()
	if messageID.ID > lastMessageID+1 || messageID.ID < 1 {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "provided non existing message ids",
		})
		c.Abort()
		return
	}
	if messageID.ID != 0 && lastMessageID+1 > messageID.ID {
		messages, err := database.GetMessagesUpFrom(messageID.ID, 100)
		messages.LastMessageID = lastMessageID
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.StatusError{
				Error: err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, messages)
		c.Abort()
		return
	}
}

// subscribe godoc
// @Summary Subscribes to updates
// @Description Subscribes to updates with longpoll connection
// @Tags Messages
// @Produce json
// @Param id query int true "Client listening message id"
// @Success 200 {object} models.MessageRes
// @Failure 400,429,500 {object} models.StatusError
// @Router /subscribe [get]
func subscribe(c *gin.Context) {
	timer := time.NewTimer(30 * time.Second)
	currentChannel := ps.Subscribe("global")
	select {
	case <-timer.C:
		c.JSON(http.StatusOK, gin.H{
			"timer": "no updates",
		})
	case <-c.Request.Context().Done():
		//nothing
	case data := <-currentChannel:
		c.JSON(http.StatusOK, data)
	}
	ps.Unsubscribe("global", currentChannel)
	close(currentChannel)
}
