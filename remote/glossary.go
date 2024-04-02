package remote

import (
	"strconv"
	"strings"

	"github.com/cufee/am-wg-proxy-next/types"
)

// Glossary
// glossary := queryPath.Group("/glossary")
// glossary.Get("/info", dummyHandlerFunc)
// glossary.Get("/achievements/:aid", dummyHandlerFunc)
// glossary.Get("/achievements", dummyHandlerFunc)
// glossary.Get("/vehicles/:vid", query.VehicleGlossaryHandler)

func (c *Client) GetOneVehicleGlossary(vehicleId int, lang string, fields ...string) (types.VehicleDetails, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("query", strconv.Itoa(vehicleId))
	opts.Query.Add("language", types.GetLocale(lang))
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target types.VehicleDetails
	return target, c.sendRequest("EU", glossaryManyVehiclesEndpoint, &target, opts)
}

// glossary.Get("/vehicles", query.AllVehiclesGlossaryHandler)
func (c *Client) GetVehiclesGlossary(lang string, fields ...string) (map[string]types.VehicleDetails, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("language", types.GetLocale(lang))
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target map[string]types.VehicleDetails
	return target, c.sendRequest("EU", glossaryManyVehiclesEndpoint, &target, opts)
}
