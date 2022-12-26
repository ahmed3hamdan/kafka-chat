package db

import (
	"context"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var Pgx *pgxpool.Pool

func init() {
	var err error
	if Pgx, err = pgxpool.New(context.Background(), config.PostgresUrl); err != nil {
		log.Fatalln(err)
	}
}
