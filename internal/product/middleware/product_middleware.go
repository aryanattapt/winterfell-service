package middleware

import (
	"fmt"
	"time"
	"wintefell-service/internal/product/model"
	"wintefell-service/internal/shared/handler"
	sharedmodel "wintefell-service/internal/shared/model"
	"wintefell-service/pkg"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateInsertProductMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Initialize the validator
		validate := validator.New()

		// Create a struct to bind the incoming JSON body
		var product model.Product_Upsert_Model
		product.ID = pkg.GenerateUUID()
		product.RegisteredDate = primitive.NewDateTimeFromTime(time.Now())
		for _, variant := range product.Variants {
			variant.ID = pkg.GenerateUUID()
			for _, subvariant := range variant.Subvariant {
				subvariant.ID = pkg.GenerateUUID()
			}
		}

		// Bind the request body to the struct
		if err := c.BodyParser(&product); err != nil {
			return handler.ErrorResponseHandler(c, sharedmodel.ErrorResponse{
				Code:       fiber.StatusBadRequest,
				Message:    "Failed to parse data",
				Stacktrace: []string{err.Error()},
				Path:       c.Path(),
				Timestamp:  time.Now().Format(time.RFC3339),
			})
		}

		// Perform validation
		var errorResultValidator []string
		if err := validate.Struct(product); err != nil {
			errorResult := pkg.ValidateForm(err)
			for _, value := range errorResult {
				errorResultValidator = append(errorResultValidator, fmt.Sprintf("%v", value))
			}
		}

		if len(errorResultValidator) > 0 {
			return handler.ErrorResponseHandler(c, sharedmodel.ErrorResponse{
				Code:       fiber.StatusBadRequest,
				Message:    "Please check your submission data",
				Stacktrace: errorResultValidator,
				Path:       c.Path(),
				Timestamp:  time.Now().Format(time.RFC3339),
			})
		}

		// Store validated product in ctx.locals for further use in handler
		c.Locals("product_payload", product)

		// Continue with the request
		return c.Next()
	}
}
