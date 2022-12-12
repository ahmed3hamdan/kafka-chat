package user

import "github.com/gofiber/fiber/v2"

func GetUserByUsername(c *fiber.Ctx) error {
	return c.SendString("user by username")
}
