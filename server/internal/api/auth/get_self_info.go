package auth

import "github.com/gofiber/fiber/v2"

func GetSelfInfo(c *fiber.Ctx) error {
	return c.SendString("selfInfo")
}
