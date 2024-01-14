package query

import (
	"github.com/cufee/am-wg-proxy-next/internal/wargaming/handlers/clans"
	"github.com/cufee/am-wg-proxy-next/types"
	"github.com/gofiber/fiber/v2"
)

func AccountClanInfoHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[*types.ClanMember]

	pid := c.Params("pid")
	realm := c.Params("realm")
	if pid == "" || realm == "" {
		response.Error.Message = "Player id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := clans.GetAccountClanInfo(realm, pid)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}

func SearchClansHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[[]types.Clan]

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
	var response types.ResponseWithError[*types.ExtendedClan]

	cid := c.Params("cid")
	realm := c.Params("realm")
	if cid == "" || realm == "" {
		response.Error.Message = "Clan id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := clans.GetClanInfo(realm, cid)
	if err != nil {
		response.Error.Message = "Nothing found"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}
