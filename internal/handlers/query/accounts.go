package query

import (
	api "github.com/byvko-dev/am-types/api/generic/v1"
	"github.com/cufee/am-wg-proxy-next/internal/wargaming/handlers/accounts"
	"github.com/gofiber/fiber/v2"
)

func SearchAccountsHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError
	query := c.Query("query", c.Query("q", ""))
	realm := c.Params("realm")
	if query == "" || realm == "" {
		response.Error.Message = "Query and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := accounts.SearchAccounts(realm, query)
	if err != nil {
		response.Error.Message = "player not found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func AccountInfoHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError

	pid := c.Params("pid")
	realm := c.Params("realm")
	if pid == "" || realm == "" {
		response.Error.Message = "Player id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := accounts.GetAccountInfo(realm, pid)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func AccountAchievementsHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError

	pid := c.Params("pid")
	realm := c.Params("realm")
	if pid == "" || realm == "" {
		response.Error.Message = "Player id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := accounts.GetAccountAchievements(realm, pid)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func AccountVehiclesHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError

	pid := c.Params("pid")
	realm := c.Params("realm")
	if pid == "" || realm == "" {
		response.Error.Message = "Player id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := accounts.GetAccountVehicles(realm, pid)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}
