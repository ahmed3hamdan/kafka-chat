package message

import "github.com/gofiber/fiber/v2"

func SendMessage(c *fiber.Ctx) error {
	return c.SendString("send a message")
}
