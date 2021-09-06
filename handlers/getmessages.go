package handlers

import (
	"chat/database"
	"chat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getMessages godoc
// @Summary Gets messages
// @Description Gets message downwards from given id with limit
// @Tags Messages
// @Produce json
// @Param last_id query int true "message id from which gets older messages"
// @Param limit query int true "limit of messages in response"
// @Success 200 {object} models.MessageRes
// @Failure 400,500 {object} models.StatusError
// @Router /messages [get]
func getMessages(c *gin.Context) {
	var messageFromLast models.MessagesFromLast
	err := c.ShouldBind(&messageFromLast)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "provided bad form query",
		})
		return
	}
	messages, err := database.GetMessagesDownFrom(messageFromLast.LastID, messageFromLast.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.StatusError{
			Error: err.Error(),
		})
		return
	}
	messages.LastMessageID = lmi.GetMessageID()
	c.JSON(http.StatusOK, messages)
}
