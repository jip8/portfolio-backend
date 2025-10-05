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

	projectsHandlers "github.com/jip/portfolio-backend/internal/api/projects/handlers"
	projectsUseCases "github.com/jip/portfolio-backend/internal/api/projects/usecases"
	projectsRepositories "github.com/jip/portfolio-backend/internal/api/projects/repositories"

	aboutHandlers "github.com/jip/portfolio-backend/internal/api/about/handlers"
	aboutUseCases "github.com/jip/portfolio-backend/internal/api/about/usecases"
	aboutRepositories "github.com/jip/portfolio-backend/internal/api/about/repositories"

	articlesHandlers "github.com/jip/portfolio-backend/internal/api/articles/handlers"
	articlesUseCases "github.com/jip/portfolio-backend/internal/api/articles/usecases"
	articlesRepositories "github.com/jip/portfolio-backend/internal/api/articles/repositories"

	contactsHandlers "github.com/jip/portfolio-backend/internal/api/contacts/handlers"
	contactsUseCases "github.com/jip/portfolio-backend/internal/api/contacts/usecases"
	contactsRepositories "github.com/jip/portfolio-backend/internal/api/contacts/repositories"
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

	// Projects
	projectsRepository := projectsRepositories.NewRepository(config, redisClient, postgresClient)
	projectsUseCase := projectsUseCases.NewUseCase(config, redisClient, projectsRepository, postgresClient)
	projectsHandler := projectsHandlers.NewHandler(projectsUseCase)
	projectsHandlers.Routes(e.Group("/projects"), projectsHandler, jwtMiddleware)

	// About
	aboutRepository := aboutRepositories.NewRepository(config, redisClient, postgresClient)
	aboutUseCase := aboutUseCases.NewUseCase(config, redisClient, aboutRepository, postgresClient)
	aboutHandler := aboutHandlers.NewHandler(aboutUseCase)
	aboutHandlers.Routes(e.Group("/about"), aboutHandler, jwtMiddleware)

	// Articles
	articlesRepository := articlesRepositories.NewRepository(config, redisClient, postgresClient)
	articlesUseCase := articlesUseCases.NewUseCase(config, redisClient, articlesRepository, postgresClient)
	articlesHandler := articlesHandlers.NewHandler(articlesUseCase)
	articlesHandlers.Routes(e.Group("/articles"), articlesHandler, jwtMiddleware)

	// Contacts
	contactsRepository := contactsRepositories.NewRepository(config, redisClient, postgresClient)
	contactsUseCase := contactsUseCases.NewUseCase(config, redisClient, contactsRepository, postgresClient)
	contactsHandler := contactsHandlers.NewHandler(contactsUseCase)
	contactsHandlers.Routes(e.Group("/contacts"), contactsHandler, jwtMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	serverPort := fmt.Sprintf(":%d", config.Server.Port)
	e.Logger.Fatal(e.Start(serverPort))
}
