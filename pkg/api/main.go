package main

import (
	"context"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/vmkevv/authservice"
	"github.com/vmkevv/authservice/actions"
	"github.com/vmkevv/authservice/ent"
	"github.com/vmkevv/authservice/mails"
	"github.com/vmkevv/authservice/services"
)

func start() error {
	// authservice.SetEnvs()
	config, err := authservice.GetConfig()
	if err != nil {
		return err
	}

	db, err := ent.Open("postgres", config.PostgresConnStr())
	if err != nil {
		return err
	}
	defer db.Close()

	context := context.Background()
	if err := db.Schema.Create(context); err != nil {
		return err
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: services.FiberErrorHandler,
	})
	appV1 := app.Group("/api/v1")

	validator := validator.New()

	userActions := actions.SetUpUser(context, db)

	emailServer := mails.NewSender(config.Mail.Host, config.Mail.Port, config.Mail.From, config.Mail.Password)

	services.SetupUserServices(userActions, validator, config, *emailServer).ServeRoutes(appV1)

	app.Listen(":8000")
	return nil
}

func main() {
	err := start()
	if err != nil {
		log.Fatalf("Error initializing app: %v", err)
	}
}
