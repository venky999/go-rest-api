package main

import (
	"context"
	"go-rest-api/internal/server"
	"log"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//zapLogger, err := zap.NewProduction()

	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	zapLogger, err := config.Build()

	if err != nil {
		log.Fatal("failed to create loggerger", err)
	}

	defer zapLogger.Sync()

	dbDSN := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(dbDSN), &gorm.Config{})

	if err != nil {
		zapLogger.Error("Error connecting to DB:", zap.Error(err))
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		zapLogger.Error("Failed to get sql.DB from gorm.DB:", zap.Error(err))
		return
	}
	defer sqlDB.Close()

	srv := server.New(db, zapLogger)

	// Start server
	go func() {
		if err := srv.Start(); err != nil {
			zapLogger.Fatal("Error connecting to DB: ", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zapLogger.Error("Error startuping the server: ", zap.Error(err))
	}
}
