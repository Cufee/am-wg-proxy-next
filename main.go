package main

import (
	"net/http"
	"os"

	"github.com/cufee/am-wg-proxy-next/internal/handlers/fast"
	"github.com/cufee/am-wg-proxy-next/internal/handlers/query"
	"github.com/cufee/am-wg-proxy-next/internal/logs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "net/http/pprof"

	pf "github.com/pkg/profile"
)

func main() {
	defer pf.Start(pf.MemProfile).Stop()
	go func() {
		err := http.ListenAndServe(":8024", nil)
		if err != nil {
			logs.Error("Failed at http server: %+v", err)
		}
	}()

	// Setup a server
	app := fiber.New()

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(logger.New())

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

	logs.Fatal("Failed to start a server: %v", app.Listen(":"+os.Getenv("PORT")))

	select {}
}

func dummyHandlerFunc(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
