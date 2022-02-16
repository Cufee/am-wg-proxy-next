package handlers

import (
	"strconv"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"aftermath.link/repo/am-wg-proxy/wargaming/handlers/clans"
	"github.com/gofiber/fiber/v2"
)

func AccountClanInfoHandler(c *fiber.Ctx) error {
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
	result, err := clans.GetAccountClanInfo(bucket, realm, playerId)
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

func SearchClansHandler(c *fiber.Ctx) error {
	var response ResponseJSON

	query := c.Query("query")
	realm := c.Params("realm")
	if query == "" || realm == "" {
		response.Error = &ResponseError{
			Message: "Query and realm are required",
		}
	}

	bucket := c.Query("bucket", client.BucketGlobal)
	result, err := clans.SearchClans(bucket, realm, query)
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

func ClanInfoHandler(c *fiber.Ctx) error {
	var response ResponseJSON

	cid := c.Params("cid")
	realm := c.Params("realm")
	if cid == "" || realm == "" {
		response.Error = &ResponseError{
			Message: "Clan id and realm are required",
		}
	}
	clanId, err := strconv.Atoi(cid)
	if err != nil {
		response.Error = &ResponseError{
			Message: "Invalid clan id",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	bucket := c.Query("bucket", client.BucketGlobal)
	result, err := clans.GetClanInfo(bucket, realm, clanId)
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
