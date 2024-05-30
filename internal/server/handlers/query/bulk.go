package query

import (
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/client"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/gofiber/fiber/v2"
)

func BulkAccountsInfoHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[map[string]types.ExtendedAccount]

		pids := c.Query("ids")
		realm := c.Params("realm")
		if pids == "" || realm == "" {
			response.Error.Message = "player id and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.BatchAccountByID(realm, strings.Split(pids, ","), strings.Split(c.Query("fields", ""), ",")...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}

func BulkAccountsAchievementsHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[map[string]types.AchievementsFrame]

		pids := c.Query("ids")
		realm := c.Params("realm")
		if pids == "" || realm == "" {
			response.Error.Message = "player id and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.BatchAccountAchievements(realm, strings.Split(pids, ","), strings.Split(c.Query("fields", ""), ",")...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}

func BulkAccountClanInfoHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[map[string]types.ClanMember]

		cids := c.Query("ids")
		realm := c.Params("realm")
		if cids == "" || realm == "" {
			response.Error.Message = "clan id and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.BatchAccountClan(realm, strings.Split(cids, ","), strings.Split(c.Query("fields", ""), ",")...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}

func BulkClanInfoHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[map[string]types.ExtendedClan]

		cids := c.Query("ids")
		realm := c.Params("realm")
		if cids == "" || realm == "" {
			response.Error.Message = "clan id and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
		result, err := wg.BatchClanByID(realm, strings.Split(cids, ","), strings.Split(c.Query("fields", ""), ",")...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}
