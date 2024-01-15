package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"

	"github.com/cufee/am-wg-proxy-next/internal/handlers/fast"
	"github.com/cufee/am-wg-proxy-next/internal/handlers/query"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"

	_ "net/http/pprof"
)

func main() {
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	// Setup a server
	app := fiber.New(fiber.Config{
		Network: os.Getenv("NETWORK"),
	})

	app.Use(fiberzerolog.New())

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
	glossary.Get("/info", dummyHandlerFunc)
	glossary.Get("/achievements/:aid", dummyHandlerFunc)
	glossary.Get("/achievements", dummyHandlerFunc)
	glossary.Get("/vehicles/:vid", query.VehicleGlossaryHandler)
	glossary.Get("/vehicles", query.AllVehiclesGlossaryHandler)

	// Bulk queries
	bulk := queryPath.Group("/bulk")
	bulk.Get("/clans/info", query.BulkClanInfoHandler)
	bulk.Get("/accounts/info", query.BulkAccountsInfoHandler)
	bulk.Get("/accounts/clan", query.BulkAccountClanInfoHandler)
	bulk.Get("/accounts/achievements", query.BulkAccountsAchievementsHandler)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

func dummyHandlerFunc(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
