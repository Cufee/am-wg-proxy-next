package workers

import (
	"fmt"

	"aftermath.link/repo/am-wg-proxy/credentials"
	"aftermath.link/repo/logs"
	"github.com/gofiber/fiber/v2"
)

func RegisterWorker(token credentials.TokenConfig) (string, error) {
	// Validate the payload
	if err := token.Validate(); err != nil {
		return "", logs.Wrap(err, "token.Validate failed")
	}

	// Validate Wargaming application ID
	ok, err := getWargamingServerStatus(token.ApplicationID)
	if err != nil {
		return "", logs.Wrap(err, "getWargamingServerStatus failed")
	}
	if !ok {
		return "", fmt.Errorf("wargaming application ID is not valid")
	}

	key, err := credentials.GenerateNewCredentialsToken(token)
	if err != nil {
		return "", logs.Wrap(err, "credentials.GenerateNewCredentialsToken failed")
	}
	return key, nil
}

func RegisterNewWorkerHandler(ctx *fiber.Ctx) error {
	// Parse the request body
	var payload credentials.TokenConfig
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Register the worker
	publicKey, err := RegisterWorker(payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "worker registered",
		"public_key": publicKey,
	})
}
