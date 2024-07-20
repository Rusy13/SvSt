package main

import (
	"WbTest/internal/infrastructure/database/postgres/database"
	"WbTest/internal/middleware"
	"WbTest/internal/mock/delivery"
	serviceOrder "WbTest/internal/mock/service"
	storageOrder "WbTest/internal/mock/storage/database"
	"WbTest/internal/routes"
	"context"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
		log.Println("Error is-----------------------", err)
	}

	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("error in logger initialization: %v", err)
	}
	logger := zapLogger.Sugar()
	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Printf("error in logger sync: %v", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := database.New(ctx)
	if err != nil {
		logger.Fatalf("error in database init: %v", err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			logger.Errorf("error in closing db")
		}
	}()

	stOrder := storageOrder.New(db, logger)

	svOrder := serviceOrder.NewMockService(stOrder)

	d := delivery.New(svOrder, logger)

	mw := middleware.New(logger)
	router := routes.GetRouter(d, mw)

	port := "8000"
	addr := ":" + port
	logger.Infow("starting server",
		"type", "START",
		"addr", addr,
	)
	logger.Fatal(http.ListenAndServe(addr, router))
}
