package service

import (
	"github.com/Ahmad940/dropify/app/model"
	"github.com/gofiber/fiber/v2"
)

var SqlNotFoundText = "record not found"

func ErrorResponse(err error, ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
		Message: err.Error(),
	})
}

func SqlErrorNotFound(err error) bool {
	return err.Error() == SqlNotFoundText
}

func SqlErrorIgnoreNotFound(err error) error {
	if err == nil {
		return nil
	}
	if err.Error() == SqlNotFoundText {
		return nil
	}
	return err
}
