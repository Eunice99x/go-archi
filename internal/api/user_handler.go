package api

import (
	"context"

	"github.com/eunice99x/dingoSuck/internal/service"
	"github.com/gofiber/fiber/v2"
)


type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	app.Post("/users", h.CreateUser)
	app.Get("/users", h.GetAllUsers)
}


func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	var body struct {
		Name string `json:"name"`
	}
		if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := u.service.CreateUser(context.Background(), body.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
} 

func (u *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := u.service.GetAllUsers(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(users)
}