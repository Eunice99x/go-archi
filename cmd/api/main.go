package main

import (
	"log"

	"github.com/eunice99x/dingoSuck/di/dic/dic"
	"github.com/eunice99x/dingoSuck/internal/api"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main(){
	builder, err := dic.NewBuilder()
	if err != nil {
		log.Fatalf("error trying to initialize the builder: %v", err.Error())
	}

	ctn := builder.Build()
	defer ctn.Delete()

	userHandler, err := ctn.SafeGet("user_handler")
	if err != nil {
		log.Fatalf("failed to get user_handler: %v", err)
	}

	app := fiber.New()

	userHandler.(*api.UserHandler).RegisterRoutes(app)
	log.Println("Server running on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}