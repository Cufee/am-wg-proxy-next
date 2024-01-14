package query

import (
	"github.com/cufee/am-wg-proxy-next/internal/wargaming/handlers/accounts"
	"github.com/cufee/am-wg-proxy-next/internal/wargaming/handlers/clans"
	"github.com/cufee/am-wg-proxy-next/types"
	"github.com/gofiber/fiber/v2"
)

func BulkAccountsInfoHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[map[string]types.ExtendedAccount]

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
	var response types.ResponseWithError[map[string]types.AchievementsFrame]

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

func BulkAccountClanInfoHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[map[string]types.ClanMember]

	cids := c.Query("ids")
	realm := c.Params("realm")
	if cids == "" || realm == "" {
		response.Error.Message = "Clan id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := clans.GetBulkAccountClanInfo(realm, cids)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func BulkClanInfoHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[map[string]types.ExtendedClan]

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
