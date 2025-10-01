package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	ctx := context.Background()

	inMemoryRepository := repository.NewInMemoryRepository()
	service := service.NewService(inMemoryRepository)

	fare := &domain.RideFareModel{
		ID:          primitive.NewObjectID(),
		UserID:      "42",
		PackageSlug: "van",
	}

	trip, err := service.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}

	log.Printf("trip created successfully: %+v\n", trip)

	// keep the program running for now
	for {
		time.Sleep(time.Second)
	}
}
