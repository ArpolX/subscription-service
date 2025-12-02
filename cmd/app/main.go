package main

import (
	"context"
	"log"
	"subscription-service/internal/app"
)

// @title subscription-service
// @description CRUDL-операции по работке с подписками.
// @version 1.0

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	ctx := context.Background()

	if err := app.RunApp(ctx); err != nil {
		log.Fatalf("error running subscription_service service: %v", err)
	}
}
