package handler

import (
	"time"
	"wintefell-service/internal/shared/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorResponseHandler(ctx *fiber.Ctx, payload model.ErrorResponse) error {
	return ctx.Status(payload.Code).JSON(payload)
}

func SuccessResponseNoDataHandler(ctx *fiber.Ctx, payload model.SuccessResponseNoData) error {
	return ctx.Status(payload.Code).JSON(payload)
}

func SuccessResponseHandler(ctx *fiber.Ctx, payload model.SuccessResponse) error {
	return ctx.Status(payload.Code).JSON(payload)
}

func MethodNotAllowedRoute(ctx *fiber.Ctx) error {
	return ErrorResponseHandler(ctx, model.ErrorResponse{
		Code:       fiber.StatusMethodNotAllowed,
		Message:    "Sorry, method is not allowed in this URL!",
		Stacktrace: []string{},
		Path:       ctx.Path(),
		Timestamp:  time.Now().Format(time.RFC3339),
	})
}

func NotFoundRoute(ctx *fiber.Ctx) error {
	return ErrorResponseHandler(ctx, model.ErrorResponse{
		Code:       fiber.StatusNotFound,
		Message:    "Sorry, destination not found",
		Stacktrace: []string{},
		Path:       ctx.Path(),
		Timestamp:  time.Now().Format(time.RFC3339),
	})
}
