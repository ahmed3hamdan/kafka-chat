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
	"strconv"
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

	messageKey := messageKeyGenerator()
	fromUserID := c.Locals("userID").(int64)
	toUserID := body.UserID
	content := body.Content

	messagesMap := make(map[int64]model.Message)

	messagesMap[fromUserID] = model.Message{
		OwnerUserID: fromUserID,
		FromUserID:  fromUserID,
		ToUserID:    toUserID,
		Key:         messageKey,
		Content:     content,
	}

	messagesMap[toUserID] = model.Message{
		OwnerUserID: toUserID,
		FromUserID:  fromUserID,
		ToUserID:    toUserID,
		Key:         messageKey,
		Content:     content,
	}

	messages := make([]model.Message, 0, len(messagesMap))
	for _, message := range messagesMap {
		messages = append(messages, message)
	}

	if err := model.InsertMessages(c.Context(), messages); err != nil {
		return err
	}

	kafkaMessages := make([]*sarama.ProducerMessage, len(messages))
	for i, message := range messages {
		kafkaKey := strconv.FormatInt(message.OwnerUserID, 10)
		kafkaValue, err := json.Marshal(kafka.Message{
			Key:        message.Key,
			FromUserID: message.FromUserID,
			ToUserID:   message.ToUserID,
			Content:    message.Content,
			CreatedAt:  time.Now(),
		})
		if err != nil {
			return err
		}
		kafkaMessages[i] = &sarama.ProducerMessage{
			Topic: "message",
			Key:   sarama.StringEncoder(kafkaKey),
			Value: sarama.ByteEncoder(kafkaValue),
		}
	}

	if err := producer.SendMessages(kafkaMessages); err != nil {
		return err
	}

	return c.JSON(api.SendMessageResponse{MessageKey: messageKey})
}
