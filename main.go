package main

import (
	"strings"

	"aftermath.link/repo/am-wg-proxy/workers"
	"aftermath.link/repo/logs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/skip"
)

type ProxyRequest struct {
	URL string `json:"url"`
}

func main() {
	// Setup a server
	app := fiber.New()

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(logger.New())
	app.Use(skip.New(dummyHandlerfunc, skipAuthCheckForIssue))

	// API Keys
	keys := app.Group("/workers")
	keys.Post("/register", workers.RegisterNewWorkerHandler)

	// Selecting a realm
	query := app.Group("/query/:realm")

	// Accounts
	accounts := query.Group("/accounts")
	accounts.Get("/info", dummyHandlerfunc)
	accounts.Get("/search", dummyHandlerfunc)
	accounts.Get("/:pid", dummyHandlerfunc)
	accounts.Get("/:pid/clan", dummyHandlerfunc)
	accounts.Get("/:pid/vehicles", dummyHandlerfunc)
	accounts.Get("/:pid/achievements", dummyHandlerfunc)

	// Clans
	clans := query.Group("/clans")
	clans.Get("/:cid", dummyHandlerfunc)
	clans.Get("/search", dummyHandlerfunc)

	// Glossary
	glossary := query.Group("/glossary")
	glossary.Get("/info", dummyHandlerfunc)
	glossary.Get("/vehicles", dummyHandlerfunc)
	glossary.Get("/vehicles/:vid", dummyHandlerfunc)
	glossary.Get("/achievements", dummyHandlerfunc)
	glossary.Get("/achievements/:aid", dummyHandlerfunc)

	logs.Fatal("Failed to start a server: %v", app.Listen(":3000"))
}

func skipAuthCheckForIssue(ctx *fiber.Ctx) bool {
	if !strings.HasSuffix(ctx.Path(), "/keys/issue") {
		return false
	}
	logs.Debug("skipAuthCheckForIssue: path matched")
	return true
}

func dummyHandlerfunc(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
