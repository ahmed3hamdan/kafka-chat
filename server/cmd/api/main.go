package main

import (
	"github.com/ahmed3hamdan/kafka-chat/server/internal/api/auth"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/api/message"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/api/user"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

func init() {
	migration, err := migrate.New("file://migrations", config.PostgresUrl)
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
		CaseSensitive: true,
	})

	app.Use(logger.New())

	apiRoute := app.Group("/api")

	authRoute := apiRoute.Group("/auth")
	userRoute := apiRoute.Group("/user")
	messageRoute := apiRoute.Group("/message")

	authRoute.Post("/login", auth.Login)
	authRoute.Post("/register", auth.Register)
	authRoute.Post("/get-self-info", auth.RequireAuthMiddleware, auth.GetSelfInfo)
	userRoute.Post("/get-by-username", auth.RequireAuthMiddleware, user.GetUserByUsername)
	messageRoute.Post("/", auth.RequireAuthMiddleware, message.SendMessage)

	app.Hooks().OnListen(func() error {
		logrus.Infoln("api is listening on " + config.ApiAddress)
		return nil
	})

	err := app.Listen(config.ApiAddress)
	if err != nil {
		logrus.Fatalln(err)
	}
}
