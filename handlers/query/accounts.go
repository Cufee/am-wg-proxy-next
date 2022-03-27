package query

import (
	"strconv"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"aftermath.link/repo/am-wg-proxy/wargaming/handlers/accounts"
	"github.com/byvko-dev/am-types/api/v1"
	"github.com/gofiber/fiber/v2"
)

func SearchAccountsHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError
	query := c.Query("query")
	realm := c.Params("realm")
	if query == "" || realm == "" {
		response.Error.Message = "Query and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	bucket := c.Query("bucket", client.BucketGlobal)
	result, err := accounts.SearchAccounts(bucket, realm, query)
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
	playerId, err := strconv.Atoi(pid)
	if err != nil {
		response.Error.Message = "Invalid player id"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	bucket := c.Query("bucket", client.BucketGlobal)
	result, err := accounts.GetAccountInfo(bucket, realm, playerId)
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
	playerId, err := strconv.Atoi(pid)
	if err != nil {
		response.Error.Message = "Invalid player id"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	bucket := c.Query("bucket", client.BucketGlobal)
	result, err := accounts.GetAccountAchievements(bucket, realm, playerId)
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
	playerId, err := strconv.Atoi(pid)
	if err != nil {
		response.Error.Message = "Invalid player id"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	bucket := c.Query("bucket", client.BucketGlobal)
	result, err := accounts.GetAccountVehicles(bucket, realm, playerId)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}
