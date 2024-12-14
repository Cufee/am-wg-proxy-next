package fast

import (
	"github.com/cufee/am-wg-proxy-next/v2/client/common"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/gofiber/fiber/v2"
)

func AccountRealmByIDHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[string]
	realm, ok := common.RealmFromID(c.Params("id"))
	if !ok {
		response.Error.Message = "invalid player id"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Data = realm.String()
	return c.JSON(response)
}
