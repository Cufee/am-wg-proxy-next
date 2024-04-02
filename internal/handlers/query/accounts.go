package query

import (
	"strings"

	"github.com/cufee/am-wg-proxy-next/accounts"
	"github.com/cufee/am-wg-proxy-next/types"
	"github.com/gofiber/fiber/v2"
)

func SearchAccountsHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[[]types.Account]
	query := c.Query("query", c.Query("q", ""))
	realm := c.Params("realm")
	if query == "" || realm == "" {
		response.Error.Message = "Query and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := accounts.SearchAccounts(realm, query, strings.Split(c.Query("fields", ""), ",")...)
	if err != nil {
		response.Error.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func AccountInfoHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[*types.ExtendedAccount]

	pid := c.Params("pid")
	realm := c.Params("realm")
	if pid == "" || realm == "" {
		response.Error.Message = "player id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := accounts.GetAccountInfo(realm, pid, strings.Split(c.Query("fields", ""), ",")...)
	if err != nil {
		response.Error.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func AccountAchievementsHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[*types.AchievementsFrame]

	pid := c.Params("pid")
	realm := c.Params("realm")
	if pid == "" || realm == "" {
		response.Error.Message = "player id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := accounts.GetAccountAchievements(realm, pid, strings.Split(c.Query("fields", ""), ",")...)
	if err != nil {
		response.Error.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func AccountVehiclesHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[[]types.VehicleStatsFrame]

	pid := c.Params("pid")
	realm := c.Params("realm")
	if pid == "" || realm == "" {
		response.Error.Message = "player id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := accounts.GetAccountVehicles(realm, pid, strings.Split(c.Query("fields", ""), ",")...)
	if err != nil {
		response.Error.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}
