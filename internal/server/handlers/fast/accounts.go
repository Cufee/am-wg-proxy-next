package fast

import (
	"github.com/cufee/am-wg-proxy-next/v2/internal/utils"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/gofiber/fiber/v2"
)

func AccountRealmByIDHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[string]
	realm := utils.RealmFromID(c.Params("id"))
	if realm == nil {
		response.Error.Message = "invalid player id"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Data = realm.String()
	return c.JSON(response)
}
