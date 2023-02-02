package message

import (
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/api"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/model"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/utils"
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
	messages, err := model.GetMessagesHistory(c.Context(), body.PageKey, ownerUserID, body.WithUserID, body.Limit+1)
	if err != nil {
		return err
	}

	var response api.GetHistoryResponse
	pageLength := utils.MinInt(len(messages), body.Limit)
	response.Entries = make([]api.Message, pageLength)
	for i := 0; i < pageLength; i++ {
		response.Entries[i] = api.Message{
			Key:        messages[i].Key,
			FromUserID: messages[i].FromUserID,
			ToUserID:   messages[i].ToUserID,
			Content:    messages[i].Content,
			CreatedAt:  messages[i].CreatedAt,
		}
	}

	if len(messages) > body.Limit {
		response.NextPageKey = &messages[body.Limit].Key
	}

	return c.JSON(response)
}
