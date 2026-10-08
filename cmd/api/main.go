package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"example/expense-api/internal/config"
	"example/expense-api/internal/database"
	"example/expense-api/internal/handler"
	"example/expense-api/internal/repository"
	"example/expense-api/internal/service"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := database.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	userRepo := repository.NewUserRepository(pool)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	r := gin.Default()
	r.GET("/healthz", handler.Health(pool))
	r.GET("/users/:id", userHandler.GetByID)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Printf("server: %v", err)
	}
}
