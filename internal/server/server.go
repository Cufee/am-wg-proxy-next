package server

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/cufee/am-wg-proxy-next/v2/client"
	"github.com/cufee/am-wg-proxy-next/v2/internal/server/handlers/fast"
	"github.com/cufee/am-wg-proxy-next/v2/internal/server/handlers/query"
	"github.com/gofiber/fiber/v2"

	_ "net/http/pprof"
)

type Options struct {
	Network string
	Port    string
}

func Listen(wgClient client.Client, opts Options, middleware ...func(c *fiber.Ctx) error) error {
	// Setup a server
	app := fiber.New(fiber.Config{
		Network: opts.Network,
	})

	for _, m := range middleware {
		app.Use(m)
	}

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	v1 := app.Group("/v1")

	// Quick checks
	fastPath := v1.Group("/fast")
	fastPath.Get("/account/id/:id/realm", fast.AccountRealmByIDHandler)

	// Selecting a realm
	queryPath := v1.Group("/query/:realm")

	// Accounts
	accounts := queryPath.Group("/accounts")
	accounts.Get("/search", query.SearchAccountsHandler(wgClient))
	accounts.Get("/:pid/achievements", query.AccountAchievementsHandler(wgClient))
	accounts.Get("/:pid/vehicles", query.AccountVehiclesHandler(wgClient))
	accounts.Get("/:pid/clan", query.AccountClanInfoHandler(wgClient))
	accounts.Get("/:pid", query.AccountInfoHandler(wgClient))

	// Clans
	clans := queryPath.Group("/clans")
	clans.Get("/search", query.SearchClansHandler(wgClient))
	clans.Get("/:cid", query.ClanInfoHandler(wgClient))

	// Glossary
	glossary := queryPath.Group("/glossary")
	glossary.Get("/info", dummyHandlerFunc)
	glossary.Get("/achievements/:aid", dummyHandlerFunc)
	glossary.Get("/achievements", dummyHandlerFunc)
	glossary.Get("/vehicles/:vid", query.VehicleGlossaryHandler(wgClient))
	glossary.Get("/vehicles", query.AllVehiclesGlossaryHandler(wgClient))

	// Bulk queries
	bulk := queryPath.Group("/bulk")
	bulk.Get("/clans/info", query.BulkClanInfoHandler(wgClient))
	bulk.Get("/accounts/info", query.BulkAccountsInfoHandler(wgClient))
	bulk.Get("/accounts/clan", query.BulkAccountClanInfoHandler(wgClient))
	bulk.Get("/accounts/achievements", query.BulkAccountsAchievementsHandler(wgClient))

	return app.Listen(":" + opts.Port)
}

func dummyHandlerFunc(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
