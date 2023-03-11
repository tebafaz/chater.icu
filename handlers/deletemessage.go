package handlers

import (
	"net/http"

	"github.com/tebafaz/chater.icu/database"
	"github.com/tebafaz/chater.icu/models"

	"github.com/gin-gonic/gin"
)

// deleteMessage godoc
// @Summary Delete message
// @Description Delete message by id
// @Tags Messages
// @Accept json
// @Produce  json
// @Param id body models.ID true "Deletes message based on message id and provided token"
// @Security ApiKeyAuth
// @Success 200 {object} models.StatusSuccess
// @Failure 400,500 {object} models.StatusError
// @Failure 401 "Unauthorized"
// @Router /user/message [delete]
func deleteMessage(c *gin.Context) {
	var messageID models.ID
	err := c.ShouldBindJSON(&messageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: "provided bad json body",
		})
		return
	}
	userInterface, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusInternalServerError, models.StatusError{
			Error: "bad context user",
		})
		return
	}
	user, ok := userInterface.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.StatusError{
			Error: "bad interface",
		})
		return
	}
	err = database.DeleteMessage(messageID.ID, user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatusError{
			Error: err.Error(),
		})
		return
	}
	deletedMessage := models.DeletedMessage{
		DeletedID: messageID.ID,
	}
	ps.Publish("global", deletedMessage)

	c.JSON(http.StatusOK, models.StatusSuccess{
		Success: "message deleted",
	})
}
