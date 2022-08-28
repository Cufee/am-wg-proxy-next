package query

import (
	"strconv"

	"aftermath.link/repo/am-wg-proxy/wargaming/handlers/clans"
	"github.com/byvko-dev/am-types/api/v1"
	"github.com/gofiber/fiber/v2"
)

func AccountClanInfoHandler(c *fiber.Ctx) error {
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

	result, err := clans.GetAccountClanInfo(realm, playerId)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func SearchClansHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError

	query := c.Query("query")
	realm := c.Params("realm")
	if query == "" || realm == "" {
		response.Error.Message = "Query and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := clans.SearchClans(realm, query)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func ClanInfoHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError

	cid := c.Params("cid")
	realm := c.Params("realm")
	if cid == "" || realm == "" {
		response.Error.Message = "Clan id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	clanId, err := strconv.Atoi(cid)
	if err != nil {
		response.Error.Message = "Invalid clan id"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := clans.GetClanInfo(realm, clanId)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}
