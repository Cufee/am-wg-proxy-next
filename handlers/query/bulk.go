package query

import (
	"aftermath.link/repo/am-wg-proxy/wargaming/handlers/accounts"
	api "github.com/byvko-dev/am-types/api/generic/v1"
	"github.com/gofiber/fiber/v2"
)

func BulkAccountsInfoHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError

	pids := c.Query("ids")
	realm := c.Params("realm")
	if pids == "" || realm == "" {
		response.Error.Message = "Player id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := accounts.GetBulkAccountsInfo(realm, pids)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func BulkAccountsVehiclesHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError

	pids := c.Query("ids")
	realm := c.Params("realm")
	if pids == "" || realm == "" {
		response.Error.Message = "Player id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := accounts.GetBulkAccountsVehicles(realm, pids)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func BulkAccountsAchievementsHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError

	pids := c.Query("ids")
	realm := c.Params("realm")
	if pids == "" || realm == "" {
		response.Error.Message = "Player id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := accounts.GetBulkAccountsAchievements(realm, pids)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}