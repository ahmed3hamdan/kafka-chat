package conversation

import (
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/api"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/model"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/utils"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func ListConversations(c *fiber.Ctx) error {
	var body api.ListConversationsRequestBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.InvalidRequestBody(err.Error()))
	}

	if err := validator.Validate.Struct(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.InvalidRequestBody(err.Error()))
	}

	ownerUserID := c.Locals("userID").(int64)
	conversations, err := model.ListConversations(c.Context(), body.PageKey, ownerUserID, body.Limit+1)
	if err != nil {
		return err
	}

	var response api.ListConversationsResponse
	pageLength := utils.MinInt(len(conversations), body.Limit)
	response.Entries = make([]api.Conversation, pageLength)
	for i := 0; i < pageLength; i++ {
		response.Entries[i] = api.Conversation{
			Key:                   conversations[i].Key,
			WithUserID:            conversations[i].WithUserID,
			WithName:              conversations[i].WithName,
			WithUsername:          conversations[i].WithUsername,
			LastMessageFromUserID: conversations[i].LastMessageFromUserID,
			LastMessageContent:    conversations[i].LastMessageContent,
			CreatedAt:             conversations[i].CreatedAt,
			UpdatedAt:             conversations[i].UpdatedAt,
		}
	}

	if len(conversations) > body.Limit {
		response.NextPageKey = &conversations[body.Limit].Key
	}

	return c.JSON(response)
}
