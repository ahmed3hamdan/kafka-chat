package model

import (
	"context"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/db"
	"github.com/jackc/pgx/v5"
	"time"
)

type Conversation struct {
	ConversationID        int64
	OwnerUserID           int64
	WithUserID            int64
	WithName              string
	WithUsername          string
	Key                   string
	LastMessageFromUserID int64
	LastMessageContent    string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

const insertConversationSql = `
	INSERT INTO "conversation" ("ownerUserID", "withUserID", "key", "lastMessageFromUserID", "lastMessageContent", "updatedAt")
	VALUES ($1, $2, $3, $4, $5, now())
	ON CONFLICT ON CONSTRAINT "conversation_ownerUserID_withUserID_key" DO UPDATE
		SET "lastMessageFromUserID" = $4,
			"lastMessageContent"    = $5,
			"updatedAt"             = now()
`

func InsertOrUpdateConversation(ctx context.Context, fromUserID int64, toUserID int64, key, content string) error {
	batch := &pgx.Batch{}
	batch.Queue(insertConversationSql, fromUserID, toUserID, key, fromUserID, content)
	if fromUserID != toUserID {
		batch.Queue(insertConversationSql, toUserID, fromUserID, key, fromUserID, content)
	}
	_, err := db.Pgx.SendBatch(ctx, batch).Exec()
	return err
}

const initialListConversationsSql = `
	SELECT "c"."conversationID",
		   "c"."ownerUserID",
		   "c"."withUserID",
		   "u"."name"     as "withName",
		   "u"."username" as "withUsername",
		   "c"."key",
		   "c"."lastMessageFromUserID",
		   "c"."lastMessageContent",
		   "c"."createdAt",
		   "c"."updatedAt"
	FROM "conversation" "c"
			 LEFT JOIN "user" "u" on "u"."userID" = "c"."lastMessageFromUserID"
	WHERE "c"."ownerUserID" = $1
	ORDER BY "c"."conversationID" DESC
	LIMIT $2
`

const listConversationsSql = `
	SELECT "c"."conversationID",
		   "c"."ownerUserID",
		   "c"."withUserID",
		   "u"."name"     as "withName",
		   "u"."username" as "withUsername",
		   "c"."key",
		   "c"."lastMessageFromUserID",
		   "c"."lastMessageContent",
		   "c"."createdAt",
		   "c"."updatedAt"
	FROM "conversation" AS "c"
			 LEFT JOIN "conversation" AS "l" ON
				"l"."key" = $1 AND "l"."ownerUserID" = "c"."ownerUserID"
			 LEFT JOIN "user" "u" on "u"."userID" = "c"."lastMessageFromUserID"
	WHERE "c"."ownerUserID" = $2
	  AND "c"."conversationID" <= "l"."conversationID"
	ORDER BY "c"."conversationID" DESC
	LIMIT $3
`

func ListConversations(ctx context.Context, conversationKey string, ownerUserID int64, limit int) ([]Conversation, error) {
	var (
		rows pgx.Rows
		err  error
	)
	if conversationKey == "" {
		rows, err = db.Pgx.Query(ctx, initialListConversationsSql, ownerUserID, limit)
	} else {
		rows, err = db.Pgx.Query(ctx, listConversationsSql, conversationKey, ownerUserID, limit)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	conversations := make([]Conversation, 0)
	for rows.Next() {
		var c Conversation
		if err = rows.Scan(&c.ConversationID,
			&c.OwnerUserID,
			&c.WithUserID,
			&c.WithName,
			&c.WithUsername,
			&c.Key,
			&c.LastMessageFromUserID,
			&c.LastMessageContent,
			&c.CreatedAt,
			&c.UpdatedAt,
		); err != nil {
			return nil, err
		}
		conversations = append(conversations, c)
	}
	return conversations, nil
}
