package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	UserName string
	Password string
}

var firstUser User

func getUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(firstUser)
}

func createUser(c *fiber.Ctx) error {
	newUser := new(User)
	err := c.BodyParser(newUser)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	firstUser = User{
		UserName: newUser.UserName,
		Password: newUser.Password,
	}

	return c.Status(fiber.StatusOK).SendString("User signed up successfully.")
}

func main() {
	app := fiber.New()

	mainApp := app.Group("/api/user")
	mainApp.Get("", getUser)
	mainApp.Post("", createUser)

	app.Listen(":8000")

}