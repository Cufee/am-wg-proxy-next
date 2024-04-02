package query

import (
	"strings"

	"github.com/cufee/am-wg-proxy-next/clans"
	"github.com/cufee/am-wg-proxy-next/types"
	"github.com/gofiber/fiber/v2"
)

func AccountClanInfoHandler(c *fiber.Ctx) error {
	var response types.ResponseWithError[*types.ClanMember]

	pid := c.Params("pid")
	realm := c.Params("realm")
	if pid == "" || realm == "" {
		response.Error.Message = "player id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := clans.GetAccountClanInfo(realm, pid, strings.Split(c.Query("fields", ""), ",")...)
	if err != nil {
		response.Error.Message = err.Error()
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
		response.Error.Message = "query and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := clans.SearchClans(realm, query, strings.Split(c.Query("fields", ""), ",")...)
	if err != nil {
		response.Error.Message = err.Error()
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
		response.Error.Message = "clan id and realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := clans.GetClanInfo(realm, cid, strings.Split(c.Query("fields", ""), ",")...)
	if err != nil {
		response.Error.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = result
	return c.JSON(response)
}
