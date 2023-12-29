package fast

import (
	"strconv"

	"github.com/cufee/am-wg-proxy-next/types"
	"github.com/gofiber/fiber/v2"
)

func AccountRealmByIDHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		response.Error.Message = "Invalid player id"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	realm := RealmFromID(id)
	if realm == "" {
		response.Error.Message = "Invalid player id"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Data = realm
	return c.JSON(response)
}

func RealmFromID(pidInt int) string {
	switch {
	case pidInt == 0:
		return ""
	case pidInt < 500000000:
		return "RU"
	case pidInt < 1000000000:
		return "EU"
	case pidInt < 2000000000:
		return "NA"
	default:
		return "AS"
	}
}
