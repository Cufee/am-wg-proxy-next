package query

import (
	"github.com/cufee/am-wg-proxy-next/internal/wargaming/handlers/glossary"
	"github.com/cufee/am-wg-proxy-next/types"
	"github.com/gofiber/fiber/v2"
)

func VehicleGlossaryHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[*types.VehicleDetails]

	vid := c.Params("id")
	realm := c.Params("realm")
	lang := c.Params("language", "en")
	if vid == "" || realm == "" {
		response.Error.Message = "Vehicle id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := glossary.GetGlossaryVehicle(realm, vid, lang)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func AllVehiclesGlossaryHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[map[string]types.VehicleDetails]

	realm := c.Params("realm")
	lang := c.Params("language", "en")
	if realm == "" {
		response.Error.Message = "realm is required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := glossary.GetAllGlossaryVehicles(realm, lang)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}
