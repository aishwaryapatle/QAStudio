package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aishwaryapatle/qastudio/internal/auth/handler"
	"github.com/aishwaryapatle/qastudio/internal/auth/repository"
	"github.com/aishwaryapatle/qastudio/internal/auth/service"
	"github.com/aishwaryapatle/qastudio/internal/config"
	"github.com/aishwaryapatle/qastudio/internal/db"
	"github.com/aishwaryapatle/qastudio/internal/logger"
	"github.com/aishwaryapatle/qastudio/internal/routes.go"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	logger.Init(cfg.AppEnv)
	logger.Log.Infof("starting server in %s mode on port %s", cfg.AppEnv, cfg.AppPort)

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	_, err := db.Connect()
	if err != nil {
		log.Fatal("DB connection failed : ",err)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	// Auth
	authService := service.NewAuthService(repository.NewUserRepo(db.DB))
	authHandler := handler.NewAuthHandler(authService)
	routes.RegisterAuthRoutes(api, authHandler)

	// Health
	routes.RegisterHealthRoutes(api)

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/api/v1/")
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.AppPort),
		Handler: r,
	}

	go func() {
		logger.Log.Infof("listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatalf("listen: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Infof("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatalf("server forced to shutdown: %v", err)
	}

	if err := db.Close(); err != nil {
		logger.Log.Errorf("error closing DB: %v", err)
	}

	logger.Log.Infof("server exiting")
}
