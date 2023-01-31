package message

import (
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/api"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/model"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func GetHistory(c *fiber.Ctx) error {
	var body api.GetHistoryRequestBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.InvalidRequestBody(err.Error()))
	}

	if err := validator.Validate.Struct(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.InvalidRequestBody(err.Error()))
	}

	ownerUserID := c.Locals("userID").(int64)
	messages, err := model.GetMessagesHistory(c.Context(), body.BeforeKey, ownerUserID, body.WithUserID, body.Limit)
	if err != nil {
		return err
	}

	l := len(messages)
	apiMessages := make([]api.Message, l)
	for i, message := range messages {
		apiMessages[i] = api.Message{
			Key:        message.Key,
			FromUserID: message.FromUserID,
			ToUserID:   message.ToUserID,
			Content:    message.Content,
			CreatedAt:  message.CreatedAt,
		}
	}

	hasMore, err := model.GetMessagesHistoryHasMore(c.Context(), messages[l-1].Key, ownerUserID, body.WithUserID)
	if err != nil {
		return err
	}

	return c.JSON(api.GetHistoryResponse{
		Messages: apiMessages,
		HasMore:  hasMore,
	})
}
