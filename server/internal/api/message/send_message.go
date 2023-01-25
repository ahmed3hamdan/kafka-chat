package message

import (
	"github.com/Shopify/sarama"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/api"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/config"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/model"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	messageKey := uuid.NewString()
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

	//kafkaMessages := make([]*sarama.ProducerMessage, 1)
	//
	//kafkaMessages[0] = &sarama.ProducerMessage{
	//	Topic: "message",
	//	Key:   sarama.ByteEncoder(strconv.FormatInt(fromUserID, 10)),
	//	Value: bytesValue,
	//}

	//if fromUserID != toUserID {
	//kafkaMessages = append(kafkaMessages, &sarama.ProducerMessage{
	//	Topic: "message",
	//	Key:   sarama.ByteEncoder(strconv.FormatInt(toUserID, 10)),
	//	Value: bytesValue,
	//})
	//}

	//if err = producer.SendMessages(kafkaMessages); err != nil {
	//	return err
	//}

	return c.JSON(api.SendMessageResponse{MessageKey: messageKey})
}
