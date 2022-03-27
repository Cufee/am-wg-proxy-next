package main

import (
	"os"

	"aftermath.link/repo/am-wg-proxy/handlers/fast"
	"aftermath.link/repo/am-wg-proxy/handlers/query"
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

	prometheus := fiberprometheus.New("am-wg-proxy")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	v1 := app.Group("/v1")

	// Quick checks
	fastPath := v1.Group("/fast")
	fastPath.Get("/account/id/:id/realm", fast.AccountRealmByIDHandler)
	fastPath.Get("/account/name/:name/id", fast.AccountIDByNameHandler)

	// Selecting a realm
	queryPath := v1.Group("/query/:realm")

	// Accounts
	accounts := queryPath.Group("/accounts")
	accounts.Get("/search", query.SearchAccountsHandler)
	accounts.Get("/:pid/achievements", query.AccountAchievementsHandler)
	accounts.Get("/:pid/vehicles", query.AccountVehiclesHandler)
	accounts.Get("/:pid/clan", query.AccountClanInfoHandler)
	accounts.Get("/:pid", query.AccountInfoHandler)

	// Clans
	clans := queryPath.Group("/clans")
	clans.Get("/search", query.SearchClansHandler)
	clans.Get("/:cid", query.ClanInfoHandler)

	// Glossary
	glossary := queryPath.Group("/glossary")
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
