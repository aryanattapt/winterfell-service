package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Helper function to generate the response structure
func generateErrorResponse(code string, message string, stacktrace []string, path string) fiber.Map {
	return fiber.Map{
		"code":       code,                            // HTTP status code indicating error or success
		"message":    message,                         // A human-readable message explaining the error
		"stacktrace": []string{},                      // An empty slice (You can add actual stack trace if needed)
		"timestamp":  time.Now().Format(time.RFC3339), // Current timestamp in ISO 8601 format
		"path":       path,                            // The API endpoint path where the error occurred
	}
}

func MethodNotAllowedRoute(ctx *fiber.Ctx) error {
	errorResponse := generateErrorResponse(
		"405",
		"Sorry, method is not allowed in this URL!",
		[]string{},
		ctx.Path(),
	)
	return ctx.Status(fiber.StatusMethodNotAllowed).JSON(errorResponse)
}

func NotFoundRoute(ctx *fiber.Ctx) error {
	errorResponse := generateErrorResponse(
		"404",
		"Sorry, destination not found",
		[]string{},
		ctx.Path(),
	)
	return ctx.Status(fiber.StatusNotFound).JSON(errorResponse)
}

func NoContentRoute(ctx *fiber.Ctx) error {
	errorResponse := generateErrorResponse(
		"204",
		"No content available",
		[]string{},
		ctx.Path(),
	)
	return ctx.Status(fiber.StatusNoContent).JSON(errorResponse)
}
