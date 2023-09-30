package handler

import (
	"github.com/Ahmad940/dropify/app/model"
	"github.com/Ahmad940/dropify/app/service"
	"github.com/Ahmad940/dropify/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	var body model.Auth
	// parsing response body
	err := ctx.BodyParser(&body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// validating the user
	errors := util.ValidateStruct(body)
	if len(errors) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// retrieving the token by passing request body
	token, err := service.Login(body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}
	return ctx.JSON(fiber.Map{
		"token": token,
	})
}

func Register(ctx *fiber.Ctx) error {
	var body model.Auth
	// parsing response body
	err := ctx.BodyParser(&body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	errors := util.ValidateStruct(body)
	if len(errors) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// retrieving the token by passing request body and registering user
	err = service.CreateAccount(body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(fiber.Map{
		"message": "Account created successful",
	})
}
