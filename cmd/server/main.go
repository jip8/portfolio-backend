package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-jwt/v4"

	loginHandlers "github.com/jip/portfolio-backend/internal/api/login/handlers"
	loginUseCases "github.com/jip/portfolio-backend/internal/api/login/usecases"

	experiencesHandlers "github.com/jip/portfolio-backend/internal/api/experiences/handlers"
	experiencesUseCases "github.com/jip/portfolio-backend/internal/api/experiences/usecases"
	experiencesRepositories "github.com/jip/portfolio-backend/internal/api/experiences/repositories"
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

	db, err := services.NewPostgresClient(config)
	if err != nil {
		log.Fatalf("Failed to create Postgres client: %s", err)
	}

	log.Println("Successfully connected to Redis, Minio, and Postgres")

	e := echo.New()

	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWT.Secret),
	})

	loginUseCase := loginUseCases.NewLoginUseCase(config, redisClient)
	loginHandler := loginHandlers.NewLoginHandler(loginUseCase)
	loginHandlers.LoginRoutes(e.Group("/login"), loginHandler)

	experiencesRepository := experiencesRepositories.NewExperiencesRepository(config, redisClient, db)
	experiencesUseCase := experiencesUseCases.NewExperiencesUseCase(config, redisClient, experiencesRepository)
	experiencesHandler := experiencesHandlers.NewExperiencesHandler(experiencesUseCase)
	experiencesHandlers.ExperiencesRoutes(e.Group("/experiences"), experiencesHandler, jwtMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	serverPort := fmt.Sprintf(":%d", config.Server.Port)
	e.Logger.Fatal(e.Start(serverPort))
}
