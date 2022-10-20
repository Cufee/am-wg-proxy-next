package client

import (
	"strconv"

	e "github.com/byvko-dev/am-types/errors/v2"
	"github.com/byvko-dev/am-types/wargaming/v2/glossary"
	"github.com/cufee/am-wg-proxy-next/helpers"
)

// Glossary
// glossary := queryPath.Group("/glossary")
// glossary.Get("/info", dummyHandlerFunc)
// glossary.Get("/achievements/:aid", dummyHandlerFunc)
// glossary.Get("/achievements", dummyHandlerFunc)
// glossary.Get("/vehicles/:vid", query.VehicleGlossaryHandler)

func (c *Client) GetOneVehicleGlossary(vehicleId int, lang string) (glossary.VehicleDetails, *e.Error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("query", strconv.Itoa(vehicleId))
	opts.Query.Add("language", helpers.GetLanguageCode(lang))

	var target glossary.VehicleDetails
	return target, c.sendRequest("EU", glossaryManyVehiclesEndpoint, &target, opts)
}

// glossary.Get("/vehicles", query.AllVehiclesGlossaryHandler)
func (c *Client) GetVehiclesGlossary(lang string) (map[string]glossary.VehicleDetails, *e.Error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("language", helpers.GetLanguageCode(lang))

	var target map[string]glossary.VehicleDetails
	return target, c.sendRequest("EU", glossaryManyVehiclesEndpoint, &target, opts)
}
