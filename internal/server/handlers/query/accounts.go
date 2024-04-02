package query

import (
	"strings"

	"github.com/cufee/am-wg-proxy-next/client"
	"github.com/cufee/am-wg-proxy-next/types"
	"github.com/gofiber/fiber/v2"
)

func SearchAccountsHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[[]types.Account]
		query := c.Query("query", c.Query("q", ""))
		realm := c.Params("realm")
		if query == "" || realm == "" {
			response.Error.Message = "Query and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.SearchAccounts(realm, query, strings.Split(c.Query("fields", ""), ",")...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}

func AccountInfoHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[*types.ExtendedAccount]

		pid := c.Params("pid")
		realm := c.Params("realm")
		if pid == "" || realm == "" {
			response.Error.Message = "player id and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.GetAccountInfo(realm, pid, strings.Split(c.Query("fields", ""), ",")...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}

func AccountAchievementsHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[*types.AchievementsFrame]

		pid := c.Params("pid")
		realm := c.Params("realm")
		if pid == "" || realm == "" {
			response.Error.Message = "player id and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.GetAccountAchievements(realm, pid, strings.Split(c.Query("fields", ""), ",")...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}

func AccountVehiclesHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[[]types.VehicleStatsFrame]

		pid := c.Params("pid")
		realm := c.Params("realm")
		if pid == "" || realm == "" {
			response.Error.Message = "player id and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.GetAccountVehicles(realm, pid, strings.Split(c.Query("fields", ""), ",")...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}
