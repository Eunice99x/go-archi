package main

import (
	"log"

	"github.com/eunice99x/fiberlearn/internal/db"
	"github.com/gofiber/fiber/v2"
)


func init(){
	if err := db.ConnectDB(); err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
}

func main(){
	defer db.DBConnection.Close()

	
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"test":"data"})
	})

	log.Fatal(app.Listen(":3000"))
}