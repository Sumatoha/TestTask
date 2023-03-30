package main

import (
	"TestTask/internal/data"
	"TestTask/internal/handlers"
	"TestTask/internal/repo"
	"TestTask/internal/service"
	"log"
)

func main() {
	db, err := data.Connect()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	cache := repo.NewInMemoryCache(1000)
	repository := repo.NewPathMongo(db)
	useCases := service.NewUseCase(repository, cache)
	handler := handlers.NewHandler(useCases)

	handler.Register()
	handler.R.Run(":9090")
}
