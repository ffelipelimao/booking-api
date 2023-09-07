package web

import (
	"net/http"

	"github.com/ffelipelimao/booking/internal/entities/place"
	"github.com/ffelipelimao/booking/internal/usecases/create_place"
	"github.com/ffelipelimao/booking/internal/web"
	"github.com/gofiber/fiber/v2"
)

var statusCodeErrHandle = map[error]int{
	place.ErrInvalidCapacity: http.StatusBadRequest,
}

type createPlacetHandler struct {
	createPlaceUseCase web.CreatePlaceUseCase
}

func NewCreatePlaceHandler(createPlaceUseCase web.CreatePlaceUseCase) createPlacetHandler {
	return createPlacetHandler{
		createPlaceUseCase: createPlaceUseCase,
	}
}

func (ch createPlacetHandler) Handle(c *fiber.Ctx) error {
	ctx := c.Context()

	var inputPlace create_place.CreatePlaceInputDTO

	if err := c.BodyParser(&inputPlace); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "invalid place body",
		})
	}

	output, err := ch.createPlaceUseCase.Create(ctx, &inputPlace)
	if err != nil {
		if statusCode, ok := statusCodeErrHandle[err]; ok {
			return c.Status(statusCode).JSON(fiber.Map{
				"code":    statusCode,
				"message": err.Error(),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(output)
}
