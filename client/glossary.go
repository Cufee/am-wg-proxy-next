package client

import (
	"strconv"

	"github.com/cufee/am-wg-proxy-next/types"
)

// Glossary
// glossary := queryPath.Group("/glossary")
// glossary.Get("/info", dummyHandlerFunc)
// glossary.Get("/achievements/:aid", dummyHandlerFunc)
// glossary.Get("/achievements", dummyHandlerFunc)
// glossary.Get("/vehicles/:vid", query.VehicleGlossaryHandler)

func (c *Client) GetOneVehicleGlossary(vehicleId int, lang string) (types.VehicleDetails, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("query", strconv.Itoa(vehicleId))
	opts.Query.Add("language", types.GetLocale(lang))

	var target types.VehicleDetails
	return target, c.sendRequest("EU", glossaryManyVehiclesEndpoint, &target, opts)
}

// glossary.Get("/vehicles", query.AllVehiclesGlossaryHandler)
func (c *Client) GetVehiclesGlossary(lang string) (map[string]types.VehicleDetails, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("language", types.GetLocale(lang))

	var target map[string]types.VehicleDetails
	return target, c.sendRequest("EU", glossaryManyVehiclesEndpoint, &target, opts)
}
