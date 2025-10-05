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

	coursesHandlers "github.com/jip/portfolio-backend/internal/api/courses/handlers"
	coursesUseCases "github.com/jip/portfolio-backend/internal/api/courses/usecases"
	coursesRepositories "github.com/jip/portfolio-backend/internal/api/courses/repositories"
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

	postgresClient, err := services.NewPostgresClient(config)
	if err != nil {
		log.Fatalf("Failed to create Postgres client: %s", err)
	}

	log.Println("Successfully connected to Redis, Minio, and Postgres")

	e := echo.New()

	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWT.Secret),
	})

	// Login
	loginUseCase := loginUseCases.NewLoginUseCase(config, redisClient)
	loginHandler := loginHandlers.NewLoginHandler(loginUseCase)
	loginHandlers.LoginRoutes(e.Group("/login"), loginHandler)

	// Experiences
	experiencesRepository := experiencesRepositories.NewRepository(config, redisClient, postgresClient)
	experiencesUseCase := experiencesUseCases.NewUseCase(config, redisClient, experiencesRepository, postgresClient)
	experiencesHandler := experiencesHandlers.NewHandler(experiencesUseCase)
	experiencesHandlers.Routes(e.Group("/experiences"), experiencesHandler, jwtMiddleware)

	// Courses
	coursesRepository := coursesRepositories.NewRepository(config, redisClient, postgresClient)
	coursesUseCase := coursesUseCases.NewUseCase(config, redisClient, coursesRepository, postgresClient)
	coursesHandler := coursesHandlers.NewHandler(coursesUseCase)
	coursesHandlers.Routes(e.Group("/courses"), coursesHandler, jwtMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	serverPort := fmt.Sprintf(":%d", config.Server.Port)
	e.Logger.Fatal(e.Start(serverPort))
}
