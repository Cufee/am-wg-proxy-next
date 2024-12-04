package query

import (
	"strconv"
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/client"
	"github.com/cufee/am-wg-proxy-next/v2/client/common"
	"github.com/cufee/am-wg-proxy-next/v2/internal/utils"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/gofiber/fiber/v2"
)

func AccountClanInfoHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[types.ClanMember]

		pid := c.Params("pid")
		realm := utils.ParseRealm(c.Params("realm"))
		if realm == nil {
			response.Error.Message = common.ErrRealmNotSupported.Error()
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
		if pid == "" {
			response.Error.Message = "player id and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.AccountClan(c.Context(), *realm, pid, strings.Split(c.Query("fields", ""), ",")...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}

func SearchClansHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[[]types.Clan]

		query := c.Query("query")
		realm := utils.ParseRealm(c.Params("realm"))
		if realm == nil {
			response.Error.Message = common.ErrRealmNotSupported.Error()
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
		if query == "" {
			response.Error.Message = "query and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
		var limit = 3
		if l, err := strconv.Atoi(c.Query("limit", "3")); err == nil && l > 1 {
			limit = l
		}

		result, err := wg.SearchClans(c.Context(), *realm, query, limit, strings.Split(c.Query("fields", ""), ",")...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}

func ClanInfoHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[types.ExtendedClan]

		cid := c.Params("cid")
		realm := utils.ParseRealm(c.Params("realm"))
		if realm == nil {
			response.Error.Message = common.ErrRealmNotSupported.Error()
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
		if cid == "" {
			response.Error.Message = "clan id and realm are required"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.ClanByID(c.Context(), *realm, cid, strings.Split(c.Query("fields", ""), ",")...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}
