package main

import (
	"os"

	"aftermath.link/repo/am-wg-proxy/handlers"
	"aftermath.link/repo/am-wg-proxy/logs"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Setup a server
	app := fiber.New()

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(logger.New())

	prometheus := fiberprometheus.New("my-service-name")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// Selecting a realm
	query := app.Group("/query/:realm")

	// Accounts
	accounts := query.Group("/accounts")
	accounts.Get("/search", handlers.SearchAccountsHandler)
	accounts.Get("/:pid/achievements", handlers.AccountAchievementsHandler)
	accounts.Get("/:pid/vehicles", handlers.AccountVehiclesHandler)
	accounts.Get("/:pid/clan", handlers.AccountClanInfoHandler)
	accounts.Get("/:pid", handlers.AccountInfoHandler)

	// Clans
	clans := query.Group("/clans")
	clans.Get("/search", handlers.SearchClansHandler)
	clans.Get("/:cid", handlers.ClanInfoHandler)

	// Glossary
	glossary := query.Group("/glossary")
	glossary.Get("/info", dummyHandlerfunc)
	glossary.Get("/achievements/:aid", dummyHandlerfunc)
	glossary.Get("/achievements", dummyHandlerfunc)
	glossary.Get("/vehicles/:vid", dummyHandlerfunc)
	glossary.Get("/vehicles", dummyHandlerfunc)

	logs.Fatal("Failed to start a server: %v", app.Listen(":"+os.Getenv("PORT")))
}

func dummyHandlerfunc(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
