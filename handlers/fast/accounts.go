package fast

import (
	"strconv"

	"aftermath.link/repo/am-wg-proxy/wargaming/helpers"
	"github.com/byvko-dev/am-types/api/v1"
	"github.com/gofiber/fiber/v2"
)

func AccountRealmByIDHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		response.Error.Message = "Invalid player id"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	realm := helpers.RealmFromID(id)
	if realm == "" {
		response.Error.Message = "Invalid player id"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Data = realm
	return c.JSON(response)
}

func AccountIDByNameHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError
	name := c.Params("name")
	if name == "" {
		response.Error.Message = "Invalid player name"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Error.Message = "Not implemented"
	return c.Status(fiber.StatusNotImplemented).JSON(response)

	// response.Data = realm
	// return c.JSON(response)
}
