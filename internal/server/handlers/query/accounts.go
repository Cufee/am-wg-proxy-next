package query

import (
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/client"
	"github.com/cufee/am-wg-proxy-next/v2/client/common"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/gofiber/fiber/v2"
)

func SearchAccountsHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[[]types.Account]
		query := c.Query("query", c.Query("q", ""))
		realm, ok := common.ParseRealm(c.Params("realm"))
		if !ok {
			response.Error.Message = common.ErrRealmNotSupported.Error()
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
		if query == "" {
			response.Error.Message = "Query is required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		var options []types.Option
		err := c.BodyParser(&options)
		if err != nil {
			response.Error.Message = "Invalid body"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.SearchAccounts(c.Context(), realm, query, options...)
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
		var response types.ResponseWithError[types.ExtendedAccount]

		pid := c.Params("pid")
		realm, ok := common.ParseRealm(c.Params("realm"))
		if !ok {
			response.Error.Message = common.ErrRealmNotSupported.Error()
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
		if pid == "" {
			response.Error.Message = "player id and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		var options []types.Option
		err := c.BodyParser(&options)
		if err != nil {
			response.Error.Message = "Invalid body"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.AccountByID(c.Context(), realm, pid, options...)
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
		var response types.ResponseWithError[types.AchievementsFrame]

		pid := c.Params("pid")
		realm, ok := common.ParseRealm(c.Params("realm"))
		if !ok {
			response.Error.Message = common.ErrRealmNotSupported.Error()
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
		if pid == "" {
			response.Error.Message = "player id and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		var options []types.Option
		err := c.BodyParser(&options)
		if err != nil {
			response.Error.Message = "Invalid body"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.AccountAchievements(c.Context(), realm, pid, options...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}

func AccountVehicleAchievementsHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[map[string]types.AchievementsFrame]

		pid := c.Params("pid")
		realm, ok := common.ParseRealm(c.Params("realm"))
		if !ok {
			response.Error.Message = common.ErrRealmNotSupported.Error()
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
		if pid == "" {
			response.Error.Message = "player id and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		var options []types.Option
		err := c.BodyParser(&options)
		if err != nil {
			response.Error.Message = "Invalid body"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.AccountVehicleAchievements(c.Context(), realm, pid, options...)
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
		realm, ok := common.ParseRealm(c.Params("realm"))
		if !ok {
			response.Error.Message = common.ErrRealmNotSupported.Error()
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
		if pid == "" {
			response.Error.Message = "player id and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		var options []types.Option
		err := c.BodyParser(&options)
		if err != nil {
			response.Error.Message = "Invalid body"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.AccountVehicles(c.Context(), realm, pid, strings.Split(c.Query("vehicles", ""), ","), options...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}
