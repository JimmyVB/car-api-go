package internal

import (
	"github.com/gofiber/fiber"
)

func SetErrorHandler(app *fiber.App) {
	app.Settings.ErrorHandler = func(ctx *fiber.Ctx, err error) {

		code := fiber.StatusInternalServerError
		var msg string

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			msg = e.Message
		}

		if msg == "" {
			msg = "cannot process the http call"
		}

		err = ctx.Status(code).JSON(internalError{
			Message: msg,
		})
	}
}

type internalError struct {
	Message string `json:"message"`
}
