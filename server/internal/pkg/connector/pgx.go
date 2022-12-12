package connector

import (
	"context"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var Pgx *pgxpool.Pool

func init() {
	var err error
	if Pgx, err = pgxpool.Connect(context.Background(), config.PostgresUrl); err != nil {
		log.Fatalln(err)
	}
}
