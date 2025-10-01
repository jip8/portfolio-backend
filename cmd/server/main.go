package main

import (
	"fmt"
	"log"
	"net/http"

	loginHandlers "github.com/jip/portfolio-backend/internal/api/login/handlers"
	loginUseCases "github.com/jip/portfolio-backend/internal/api/login/usecases"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/labstack/echo/v4"
)

func main() {
	config, err := entity.InitConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	redisClient, err := services.NewRedisClient(config)
	if err != nil {
		log.Fatalf("Failed to create Redis client: %s", err)
	}

	_, err = services.NewMinioClient(config)
	if err != nil {
		log.Fatalf("Failed to create Minio client: %s", err)
	}

	_, err = services.NewPostgresClient(config)
	if err != nil {
		log.Fatalf("Failed to create Postgres client: %s", err)
	}

	log.Println("Successfully connected to Redis, Minio, and Postgres")

	e := echo.New()

	loginUseCase := loginUseCases.NewLoginUseCase(config, redisClient)
	loginHandler := loginHandlers.NewLoginHandler(loginUseCase)
	loginHandlers.LoginRoutes(e.Group("/login"), loginHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	serverPort := fmt.Sprintf(":%d", config.Server.Port)
	e.Logger.Fatal(e.Start(serverPort))
}
