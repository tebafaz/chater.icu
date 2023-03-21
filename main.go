package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tebafaz/chater.icu/database"
	_ "github.com/tebafaz/chater.icu/docs"
	"github.com/tebafaz/chater.icu/handlers"
	"github.com/tebafaz/chater.icu/middlewares"
	"github.com/tebafaz/chater.icu/redis"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	err = redis.InitRedis()
	if err != nil {
		panic(err)
	}
	err = database.InitPostgres()
	if err != nil {
		panic(err)
	}
	middlewares.NewSessionCounter()
}

// @title Chater.icu API
// @version 1.0
// @description Chater api made by Tebafaz using long poll as connction

// @contact.name Mukhamedjanov Erjan
// @contact.email tebafaz@gmail.com

// @host chater.icu
// @BasePath /api/v1

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @authorizationurl https://chater.icu/api/v1/register
// @tokenUrl https://chater.icu/api/v1/login

// @schemes https

func startServer() (*http.Server, <-chan error) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(middlewares.AccessLog(false))
	router.Use(middlewares.SessionCounter(3000))

	swaggerURL := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	router.Use(gin.Recovery())

	returnChannel := make(chan error, 1)

	err := handlers.MapRoutes(router)
	if err != nil {
		returnChannel <- err
		return nil, returnChannel
	}

	srv := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			returnChannel <- err
		}
	}()
	return srv, returnChannel
}

func main() {
	server, startServerError := startServer()

	defer database.ClosePostgres()
	defer redis.CloseRedis()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Server started")
	select {
	case sig := <-quit:
		fmt.Printf("Received signal: %s\nStopping server...\n", sig.String())
		handlers.ClosePS()
	case err := <-startServerError:
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	fmt.Println("Server stopped")
}
