package model

import (
	"context"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/db"
	"github.com/jackc/pgx/v5"
	"time"
)

type Message struct {
	MessageID   int64
	OwnerUserID int64
	WithUserID  int64
	FromUserID  int64
	ToUserID    int64
	Key         string
	Content     string
	CreatedAt   time.Time
}

func InsertMessages(ctx context.Context, messages []Message) error {
	_, err := db.Pgx.CopyFrom(
		ctx, pgx.Identifier{"message"},
		[]string{"ownerUserID", "withUserID", "fromUserID", "toUserID", "key", "content"},
		pgx.CopyFromSlice(len(messages), func(i int) ([]interface{}, error) {
			return []interface{}{messages[i].OwnerUserID, messages[i].WithUserID, messages[i].FromUserID, messages[i].ToUserID, messages[i].Key, messages[i].Content}, nil
		}),
	)
	return err
}

const initialGetMessagesHistorySql = `
	SELECT "m"."messageID",
		   "m"."ownerUserID",
		   "m"."fromUserID",
		   "m"."toUserID",
		   "m"."key",
		   "m"."content",
		   "m"."createdAt"
	FROM "message" AS "m"
	WHERE "ownerUserID" = $1
	  AND "withUserID" = $2
	ORDER BY "m"."messageID" DESC
	LIMIT $3
`

const getMessagesHistorySql = `
	SELECT "m"."messageID",
		   "m"."ownerUserID",
		   "m"."fromUserID",
		   "m"."toUserID",
		   "m"."key",
		   "m"."content",
		   "m"."createdAt"
	FROM "message" AS "m"
			 LEFT JOIN "message" AS "l" ON "l"."key" = $1 AND "l"."ownerUserID" = "m"."ownerUserID"
	WHERE "m"."ownerUserID" = $2
	  AND "m"."withUserID" = $3
	  AND "m"."messageID" <= "l"."messageID"
	ORDER BY "m"."messageID" DESC
	LIMIT $4
`

func GetMessagesHistory(ctx context.Context, messageKey string, ownerUserID, withUserID int64, limit int) ([]Message, error) {
	var (
		rows pgx.Rows
		err  error
	)
	if messageKey == "" {
		rows, err = db.Pgx.Query(ctx, initialGetMessagesHistorySql, ownerUserID, withUserID, limit)
	} else {
		rows, err = db.Pgx.Query(ctx, getMessagesHistorySql, messageKey, ownerUserID, withUserID, limit)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	messages := make([]Message, 0)
	for rows.Next() {
		var m Message
		if err = rows.Scan(
			&m.MessageID,
			&m.OwnerUserID,
			&m.FromUserID,
			&m.ToUserID,
			&m.Key,
			&m.Content,
			&m.CreatedAt,
		); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	return messages, nil
}
