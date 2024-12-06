package api

import (
	"context"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

// Glossary
// glossary := queryPath.Group("/glossary")
// glossary.Get("/info", dummyHandlerFunc)
// glossary.Get("/achievements/:aid", dummyHandlerFunc)
// glossary.Get("/achievements", dummyHandlerFunc)
// glossary.Get("/vehicles/:vid", query.VehicleGlossaryHandler)

func (c *Client) VehicleGlossary(ctx context.Context, realm types.Realm, vehicleId string, options ...types.Option) (types.VehicleDetails, error) {
	opts := newDefaultRequestOptions(options)
	opts.Query.Add("query", vehicleId)
	var target types.VehicleDetails
	return target, c.sendRequest(ctx, realm, glossaryManyVehiclesEndpoint, &target, opts)
}

// glossary.Get("/vehicles", query.AllVehiclesGlossaryHandler)
func (c *Client) CompleteVehicleGlossary(ctx context.Context, realm types.Realm, options ...types.Option) (map[string]types.VehicleDetails, error) {
	opts := newDefaultRequestOptions(options)
	var target map[string]types.VehicleDetails
	return target, c.sendRequest(ctx, realm, glossaryManyVehiclesEndpoint, &target, opts)
}
