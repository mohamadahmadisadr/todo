package main

import (
	"context"
	"log"
	"time"
	"todo/config"
	"todo/helper"
	"todo/internal/model"
	"todo/internal/repository"
)

func main() {
	helper.LoadEnv()

	client := config.ConnectDB()

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
		log.Println("Connection to MongoDB closed.")
	}()

	const dbName = "app_db"
	const collectionName = "users"

	userCollection := config.GetCollection(client, dbName, collectionName)

	userRepo := repository.NewUserRepository(userCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	newUser := &model.User{
		Name:     "Alice Smith",
		Email:    "alice@example.com",
		UserType: "Admin",
	}
	insertedID, err := userRepo.CreateUser(ctx, newUser)

	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
	log.Printf("Successfully created user with ID: %s", insertedID.Hex())
}
