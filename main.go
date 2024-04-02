package main

import (
	"os"
	"strconv"

	"github.com/cufee/am-wg-proxy-next/v2/client"
	"github.com/cufee/am-wg-proxy-next/v2/internal/server"
	"github.com/cufee/am-wg-proxy-next/v2/internal/utils"
	"github.com/gofiber/contrib/fiberzerolog"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"

	_ "net/http/pprof"
)

func main() {
	level, _ := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(level)

	primaryAppID := utils.MustGetEnv("PRIMARY_WG_APP_ID")
	primaryRps, _ := strconv.Atoi(utils.MustGetEnv("PRIMARY_WG_APP_RPS"))

	buckets := client.ParseProxyString(os.Getenv("PROXY_HOST_LIST"), primaryAppID, primaryRps)
	wgClient, err := client.NewClient(primaryAppID, primaryRps, client.Options{Buckets: buckets})
	if err != nil {
		panic(err)
	}

	panic(server.Listen(wgClient, server.Options{Port: os.Getenv("PORT"), Network: os.Getenv("NETWORK")}, fiberzerolog.New()))
}
