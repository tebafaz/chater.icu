package handlers

import (
	"chat/helpers"
	"chat/middlewares"
	"errors"

	//"chat/middlewares"

	//"musicapi/middlewares"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var ps *helpers.Pubsub
var lmi *helpers.LastMessageID

func ClosePS() {
	ps.Publish("global", nil)
}

// MapRoutes Создает маршруты
func MapRoutes(router *gin.Engine) error {
	if router == nil {
		return errors.New("o my god, router not created")
	}

	ps = helpers.NewPubsub()
	var err error
	lmi, err = helpers.NewLastMessageID()

	if err != nil {
		return err
	}

	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	apiGroup := router.Group("/api")
	{
		v1 := apiGroup.Group("/v1")
		{
			v1.GET("/last-messages", getLastMessages)
			v1.GET("/subscribe", getIfExist, middlewares.CheckSessionNumber, subscribe)
			v1.GET("/messages", getMessages)
			v1.POST("/login", login)
			v1.POST("/register", registerAndLogin)
			guestGroup := v1.Group("/guest")
			{
				guestGroup.POST("/message", publishAsGuest)
			}
			userGroup := v1.Group("/user", middlewares.CheckToken)
			{
				userGroup.POST("/message", publishAsUser)
				userGroup.DELETE("/message", deleteMessage)
				userGroup.POST("/logout", logout)
			}
		}
	}
	return nil
}
