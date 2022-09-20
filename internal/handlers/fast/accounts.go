package fast

import (
	"fmt"
	"strconv"

	api "github.com/byvko-dev/am-types/api/generic/v1"
	"github.com/cufee/am-wg-proxy-next/internal/shims"
	"github.com/cufee/am-wg-proxy-next/internal/wargaming/helpers"
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
	realm := c.Params("realm")
	if name == "" || realm == "" {
		response.Error.Message = "Invalid player name or realm"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	user, err := shims.CheckUserByName(name, realm)
	if err != nil {
		response.Error.Message = "Error while checking user"
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = fmt.Sprint(user.DefaultPID)
	return c.JSON(response)
}
