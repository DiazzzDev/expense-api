package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"example/expense-api/internal/config"
	"example/expense-api/internal/database"
	"example/expense-api/internal/handler"
	"example/expense-api/internal/middleware"
	"example/expense-api/internal/repository"
	"example/expense-api/internal/service"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := database.Connect(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	auth, err := middleware.NewAuthenticator(cfg.CognitoRegion, cfg.CognitoUserPoolID, cfg.CognitoClientID)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(pool)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	r := gin.Default()
	r.GET("/healthz", handler.Health(pool))

	api := r.Group("/api/v1", auth.Require(), middleware.EnsureUser(userSvc))
	api.GET("/me", userHandler.Me)
	api.PATCH("/me", userHandler.UpdateMe)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Printf("server: %v", err)
	}
}
