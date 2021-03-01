package services

import (
	"bytes"
	"errors"
	"fmt"
	"net/mail"
	"text/template"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	auth "github.com/vmkevv/authservice"
	defs "github.com/vmkevv/authservice/definitions"
	"github.com/vmkevv/authservice/ent"
	"github.com/vmkevv/authservice/mails"
)

// UserServices contains all user rest services
type UserServices struct {
	actions   defs.UserActions
	validator *validator.Validate
	conf      auth.Config
	sender    mails.Sender
}

// SetupUserServices create a new instace of UserServices
func SetupUserServices(actions defs.UserActions, validator *validator.Validate, config auth.Config, sender mails.Sender) UserServices {
	return UserServices{
		actions:   actions,
		validator: validator,
		conf:      config,
		sender:    sender,
	}
}

// ServeRoutes serve the routes defined
func (us UserServices) ServeRoutes(app fiber.Router) {
	app.Post("/user", WithAuth([]int8{auth.AdminRole}, us.conf.App.JWTKey), us.register())
	app.Post("/email", us.sendEmail())
	app.Get("/authenticate", WithAuth([]int8{auth.AdminRole, auth.LocutorRole, auth.UserRole}, us.conf.App.JWTKey), us.authenticate())
}

func (us UserServices) register() fiber.Handler {
	type request struct {
		Name     string `json:"name" validate:"required,gte=2"`
		LastName string `json:"lastName" validate:"required,gte=2"`
		Email    string `json:"email" validate:"required,email"`
	}
	type response struct {
		User *ent.User `json:"user"`
	}
	return func(c *fiber.Ctx) error {
		reqData := request{}

		if err := c.BodyParser(&reqData); err != nil {
			return New500(err)
		}

		err := us.validator.Struct(reqData)
		if err != nil {
			return NewErr(BadFields, fmt.Sprintf("Error in provided data: %s", err.Error()), fiber.StatusBadRequest)
		}

		exists, err := us.actions.ExistEmail(reqData.Email)
		if err != nil {
			return New500(err)
		}
		if exists {
			return NewErr(EmailExists, "There is already an account with this email.", fiber.StatusConflict)
		}

		savedUser, err := us.actions.Register(reqData.Name, reqData.LastName, reqData.Email)
		if err != nil {
			return New500(err)
		}

		return c.JSON(response{
			User: savedUser,
		})
	}
}

func (us UserServices) sendEmail() fiber.Handler {
	type request struct {
		Email string `json:"email" validate:"required,email"`
	}
	type response struct {
		Message string `json:"message"`
	}
	return func(c *fiber.Ctx) error {
		reqData := request{}
		if err := c.BodyParser(&reqData); err != nil {
			return New500(err)
		}

		exists, err := us.actions.ExistEmail(reqData.Email)
		if err != nil {
			return New500(err)
		}
		if !exists {
			return NewErr(EmailNotExists, "The provided email does not exists in database", fiber.StatusBadRequest)
		}

		user, err := us.actions.GetUserByEmail(reqData.Email)
		if err != nil {
			return New500(err)
		}

		token, err := GenerateToken(us.conf.App.JWTKey, user.ID, user.Role, 15*time.Minute)

		var body bytes.Buffer
		data := map[string]string{
			"MagicLink": us.conf.App.Domain + "/token/" + token,
			"Minutes":   "15",
		}

		emailTmpl := template.Must(template.ParseFiles("templates/email.html"))
		if err := emailTmpl.Execute(&body, data); err != nil {
			return New500(err)
		}

		if err := us.sender.Send(mails.Mail{
			To:      mail.Address{Address: user.Email},
			Subject: "Enlace de Ingreso",
			Body:    body.String(),
		}); err != nil {
			return New500(err)
		}

		return c.JSON(response{
			Message: "Magic link have been send to email address",
		})
	}
}

func (us UserServices) authenticate() fiber.Handler {
	type response struct {
		User  *ent.User `json:"user"`
		Token string    `json:"token"`
	}
	return func(c *fiber.Ctx) error {
		userIDStr := c.Locals("userID")

		userID, ok := userIDStr.(int)
		if !ok {
			return New500(errors.New("Can not convert userID, to int"))
		}

		user, err := us.actions.GetUserByID(userID)
		if err != nil {
			return New500(err)
		}

		token, err := GenerateToken(us.conf.App.JWTKey, user.ID, user.Role, time.Minute*time.Duration(us.conf.App.JWTExpTime))
		if err != nil {
			return New500(err)
		}

		return c.JSON(response{
			User:  user,
			Token: token,
		})
	}
}
