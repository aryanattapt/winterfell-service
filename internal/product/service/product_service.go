package service

import (
	"time"
	"wintefell-service/internal/product/model"
	"wintefell-service/internal/product/repository"
	"wintefell-service/internal/shared/handler"
	sharedmodel "wintefell-service/internal/shared/model"

	"github.com/gofiber/fiber/v2"
)

func InsertProduct(c *fiber.Ctx) error {
	// Retrieve the product from ctx.locals
	payload, ok := c.Locals("product_payload").(model.Product_Upsert_Model)
	if !ok {
		return handler.ErrorResponseHandler(c, sharedmodel.ErrorResponse{
			Code:       fiber.StatusBadRequest,
			Message:    "Failed to parse data",
			Stacktrace: []string{},
			Path:       c.Path(),
			Timestamp:  time.Now().Format(time.RFC3339),
		})
	}

	resultChan := make(chan error, 2) // Buffer size of 2 to handle both Mongo and Elasticsearch

	// Run MongoDB and Elasticsearch insertion concurrently
	go repository.InsertProductToMongoDB(payload, resultChan)
	go repository.InsertProductToElasticSearch(payload, resultChan)

	// Collect errors
	var mongoErr, esErr error
	for i := 0; i < 2; i++ {
		err := <-resultChan
		if err != nil {
			if i == 0 {
				mongoErr = err
			} else {
				esErr = err
			}
		}
	}

	// If there were any errors, return an error response
	if mongoErr != nil || esErr != nil {
		return handler.ErrorResponseHandler(c, sharedmodel.ErrorResponse{
			Code:       fiber.StatusBadRequest,
			Message:    "Failed to insert data",
			Stacktrace: []string{mongoErr.Error(), esErr.Error()},
			Path:       c.Path(),
			Timestamp:  time.Now().Format(time.RFC3339),
		})
	}

	// Success response
	return handler.SuccessResponseNoDataHandler(c, sharedmodel.SuccessResponseNoData{
		Code:      fiber.StatusCreated,
		Message:   "Success Insert Product",
		Timestamp: time.Now().Format(time.RFC3339),
	})

}
