package main

import (
	"github.com/ahmed3hamdan/kafka-chat/server/internal/api/handlers/auth"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/api/handlers/message"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/api/handlers/user"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

func init() {
	migration, err := migrate.New("file://internal/api/migrations", config.PostgresUrl)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		logrus.Fatalln(err)
	}

	logrus.Infoln("migrations are up to date")
}

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	apiRoute := app.Group("/api")

	authRoute := apiRoute.Group("/auth")
	userRoute := apiRoute.Group("/user")
	messageRoute := apiRoute.Group("/message")

	authRoute.Post("/login", auth.Login)
	authRoute.Post("/register", auth.Register)
	authRoute.Get("/self-info", auth.GetSelfInfo)
	userRoute.Get("/:userID", user.GetUserByUsername)
	messageRoute.Post("/", message.SendMessage)

	app.Hooks().OnListen(func() error {
		logrus.Infoln("api is listening on " + config.ApiAddress)
		return nil
	})

	err := app.Listen(config.ApiAddress)
	if err != nil {
		logrus.Fatalln(err)
	}
}