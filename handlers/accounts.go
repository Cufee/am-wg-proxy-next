package handlers

import (
	"strconv"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"aftermath.link/repo/am-wg-proxy/wargaming/handlers/accounts"
	"github.com/gofiber/fiber/v2"
)

func SearchAccountsHandler(c *fiber.Ctx) error {
	var response ResponseJSON
	query := c.Query("query")
	realm := c.Params("realm")
	if query == "" || realm == "" {
		response.Error = &ResponseError{
			Message: "Query and realm are required",
		}
	}

	bucket := c.Query("bucket", client.BucketGlobal)
	result, err := accounts.SearchAccounts(bucket, realm, query)
	if err != nil {
		response.Error = &ResponseError{
			Message: "Nothing found",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func AccountInfoHandler(c *fiber.Ctx) error {
	var response ResponseJSON

	pid := c.Params("pid")
	realm := c.Params("realm")
	if pid == "" || realm == "" {
		response.Error = &ResponseError{
			Message: "Player id and realm are required",
		}
	}
	playerId, err := strconv.Atoi(pid)
	if err != nil {
		response.Error = &ResponseError{
			Message: "Invalid player id",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	bucket := c.Query("bucket", client.BucketGlobal)
	result, err := accounts.GetAccountInfo(bucket, realm, playerId)
	if err != nil {
		response.Error = &ResponseError{
			Message: "Nothing found",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func AccountAchievementsHandler(c *fiber.Ctx) error {
	var response ResponseJSON

	pid := c.Params("pid")
	realm := c.Params("realm")
	if pid == "" || realm == "" {
		response.Error = &ResponseError{
			Message: "Player id and realm are required",
		}
	}
	playerId, err := strconv.Atoi(pid)
	if err != nil {
		response.Error = &ResponseError{
			Message: "Invalid player id",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	bucket := c.Query("bucket", client.BucketGlobal)
	result, err := accounts.GetAccountAchievements(bucket, realm, playerId)
	if err != nil {
		response.Error = &ResponseError{
			Message: "Nothing found",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func AccountVehiclesHandler(c *fiber.Ctx) error {
	var response ResponseJSON

	pid := c.Params("pid")
	realm := c.Params("realm")
	if pid == "" || realm == "" {
		response.Error = &ResponseError{
			Message: "Player id and realm are required",
		}
	}
	playerId, err := strconv.Atoi(pid)
	if err != nil {
		response.Error = &ResponseError{
			Message: "Invalid player id",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	bucket := c.Query("bucket", client.BucketGlobal)
	result, err := accounts.GetAccountVehicles(bucket, realm, playerId)
	if err != nil {
		response.Error = &ResponseError{
			Message: "Nothing found",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}
