package query

import (
	api "github.com/byvko-dev/am-types/api/generic/v1"
	"github.com/cufee/am-wg-proxy-next/internal/wargaming/handlers/accounts"
	"github.com/cufee/am-wg-proxy-next/internal/wargaming/handlers/clans"
	"github.com/cufee/am-wg-proxy-next/internal/wargaming/handlers/glossary"
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

func BulkClanInfoHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError

	cids := c.Query("ids")
	realm := c.Params("realm")
	if cids == "" || realm == "" {
		response.Error.Message = "Clan id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	result, err := clans.GetBulkClanInfo(realm, cids)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func AllVehiclesGlossaryHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError

	realm := c.Params("realm")
	if realm == "" {
		response.Error.Message = "realm is required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := glossary.GetAllGlossaryVehicles(realm)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}
