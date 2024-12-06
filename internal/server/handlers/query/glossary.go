package query

import (
	"github.com/cufee/am-wg-proxy-next/v2/client"
	"github.com/cufee/am-wg-proxy-next/v2/client/common"
	"github.com/cufee/am-wg-proxy-next/v2/internal/utils"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/gofiber/fiber/v2"
)

func VehicleGlossaryHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[types.VehicleDetails]

		vid := c.Params("id")
		realm := utils.ParseRealm(c.Params("realm"))
		if realm == nil {
			response.Error.Message = common.ErrRealmNotSupported.Error()
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		var options []types.Option
		err := c.BodyParser(&options)
		if err != nil {
			response.Error.Message = "Invalid body"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.VehicleGlossary(c.Context(), *realm, vid, options...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}

func AllVehiclesGlossaryHandler(wg client.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var response types.ResponseWithError[map[string]types.VehicleDetails]

		realm := utils.ParseRealm(c.Params("realm"))
		if realm == nil {
			response.Error.Message = common.ErrRealmNotSupported.Error()
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		var options []types.Option
		err := c.BodyParser(&options)
		if err != nil {
			response.Error.Message = "Invalid body"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		result, err := wg.CompleteVehicleGlossary(c.Context(), *realm, options...)
		if err != nil {
			response.Error.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		response.Data = result
		return c.JSON(response)
	}
}
