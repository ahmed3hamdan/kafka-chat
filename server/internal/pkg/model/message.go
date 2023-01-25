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
	FromUserID  int64
	ToUserID    int64
	Key         string
	Content     string
	CreatedAt   time.Time
}

func InsertMessages(ctx context.Context, messages []Message) error {
	_, err := db.Pgx.CopyFrom(
		ctx, pgx.Identifier{"message"},
		[]string{"ownerUserID", "fromUserID", "toUserID", "key", "content"},
		pgx.CopyFromSlice(len(messages), func(i int) ([]interface{}, error) {
			return []interface{}{messages[i].OwnerUserID, messages[i].FromUserID, messages[i].ToUserID, messages[i].Key, messages[i].Content}, nil
		}),
	)
	return err
}
