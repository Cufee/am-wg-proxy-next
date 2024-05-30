package api

import (
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

// Glossary
// glossary := queryPath.Group("/glossary")
// glossary.Get("/info", dummyHandlerFunc)
// glossary.Get("/achievements/:aid", dummyHandlerFunc)
// glossary.Get("/achievements", dummyHandlerFunc)
// glossary.Get("/vehicles/:vid", query.VehicleGlossaryHandler)

func (c *Client) VehicleGlossary(realm string, vehicleId string, lang string, fields ...string) (types.VehicleDetails, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("query", vehicleId)
	opts.Query.Add("language", types.GetLocale(lang))
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target types.VehicleDetails
	return target, c.sendRequest(realm, glossaryManyVehiclesEndpoint, &target, opts)
}

// glossary.Get("/vehicles", query.AllVehiclesGlossaryHandler)
func (c *Client) CompleteVehicleGlossary(realm string, lang string, fields ...string) (map[string]types.VehicleDetails, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("language", types.GetLocale(lang))
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target map[string]types.VehicleDetails
	return target, c.sendRequest(realm, glossaryManyVehiclesEndpoint, &target, opts)
}
