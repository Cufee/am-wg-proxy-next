package query

import (
	"aftermath.link/repo/am-wg-proxy/wargaming/handlers/glossary"
	api "github.com/byvko-dev/am-types/api/generic/v1"
	"github.com/gofiber/fiber/v2"
)

func VehicleGlossaryHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError

	vid := c.Params("pid")
	realm := c.Params("realm")
	if vid == "" || realm == "" {
		response.Error.Message = "Player id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := glossary.GetGlossaryVehicle(realm, vid)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}
