package message

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/api"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/config"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/kafka"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"log"
)

var producer sarama.SyncProducer

func init() {
	var err error
	if producer, err = sarama.NewSyncProducer([]string{config.KafkaAddress}, nil); err != nil {
		log.Fatalln(err)
	}
}

func SendMessage(c *fiber.Ctx) error {
	var body api.SendMessageRequestBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.InvalidRequestBody(err.Error()))
	}

	if err := validator.Validate.Struct(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.InvalidRequestBody(err.Error()))
	}

	sendID := c.Locals("userID").(int64)

	message := kafka.Message{
		SendID:    sendID,
		ReceiveID: body.UserID,
		Content:   body.Content,
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	producerMessage := &sarama.ProducerMessage{
		Topic: "message",
		Value: sarama.ByteEncoder(jsonMessage),
	}

	_, _, err = producer.SendMessage(producerMessage)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
