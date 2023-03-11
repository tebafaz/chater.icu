package handlers

import (
	"net/http"

	"github.com/tebafaz/chater.icu/database"
	"github.com/tebafaz/chater.icu/models"

	"github.com/gin-gonic/gin"
)

// getLastMessages godoc
// @Summary Gets last messages
// @Description Gets last 100 messages and latest message id for longpoll listening
// @Tags Messages
// @Produce json
// @Success 200 {object} models.MessageRes
// @Failure 500 {object} models.StatusError
// @Router /last-messages [get]
func getLastMessages(c *gin.Context) {
	messages, err := database.GetLastMessages(100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.StatusError{
			Error: err.Error(),
		})
		return
	}
	messages.LastMessageID = lmi.GetMessageID()
	c.JSON(http.StatusOK, messages)
}
