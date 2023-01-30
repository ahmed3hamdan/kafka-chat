package message

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/api"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/config"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/kafka"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/model"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jaevor/go-nanoid"
	"github.com/sirupsen/logrus"
	"time"
)

var producer sarama.SyncProducer

func init() {
	var err error
	if producer, err = sarama.NewSyncProducer([]string{config.KafkaAddress}, nil); err != nil {
		logrus.Fatalln(err)
	}
}

var messageKeyGenerator func() string

func init() {
	var err error
	messageKeyGenerator, err = nanoid.Standard(21)
	if err != nil {
		logrus.Fatalln(err)
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

	if _, err := model.GetUserById(c.Context(), body.UserID); err == model.UserNotFoundError {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.UserNotFound(err.Error()))
	} else if err != nil {
		return err
	}

	messageKey := messageKeyGenerator()
	fromUserID := c.Locals("userID").(int64)
	toUserID := body.UserID

	kafkaValue, err := json.Marshal(kafka.MessageBody{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Key:        messageKey,
		Content:    body.Content,
		CreatedAt:  time.Now(),
	})
	if err != nil {
		return err
	}

	if _, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: "message",
		Value: sarama.ByteEncoder(kafkaValue),
	}); err != nil {
		return err
	}

	return c.JSON(api.SendMessageResponse{MessageKey: messageKey})
}
