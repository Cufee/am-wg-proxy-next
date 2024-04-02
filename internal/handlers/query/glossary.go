package query

import (
	"strings"

	"github.com/cufee/am-wg-proxy-next/glossary"
	"github.com/cufee/am-wg-proxy-next/types"
	"github.com/gofiber/fiber/v2"
)

func VehicleGlossaryHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[*types.VehicleDetails]

	vid := c.Params("id")
	realm := c.Params("realm")
	lang := c.Params("language", "en")
	if vid == "" || realm == "" {
		response.Error.Message = "vehicle id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := glossary.GetGlossaryVehicle(realm, vid, lang, strings.Split(c.Query("fields", ""), ",")...)
	if err != nil {
		response.Error.Message = err.Error()
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

	result, err := glossary.GetAllGlossaryVehicles(realm, lang, strings.Split(c.Query("fields", ""), ",")...)
	if err != nil {
		response.Error.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}
