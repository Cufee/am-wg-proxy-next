package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

func (c *Client) VehicleGlossary(ctx context.Context, realm types.Realm, vehicleID string, opts ...types.Option) (types.VehicleDetails, error) {
	var response types.WgResponse[map[string]types.VehicleDetails]
	options := types.GetOptions(opts...)
	options.Fields = append(options.Fields, "tank_id", "name", "nation", "tier", "type", "is_premium")
	query := options.Query()
	query.Set("tank_id", vehicleID)

	_, err := c.Request(ctx, realm, fmt.Sprintf("encyclopedia/vehicles/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return types.VehicleDetails{}, err
	}
	if response.Error.Code != 0 {
		return types.VehicleDetails{}, errors.New(response.Error.Message)
	}

	info, ok := response.Data[fmt.Sprint(vehicleID)]
	if !ok || info.TankID == 0 {
		return types.VehicleDetails{}, errors.New("vehicle not found")
	}
	return info, nil
}

func (c *Client) CompleteVehicleGlossary(ctx context.Context, realm types.Realm, opts ...types.Option) (map[string]types.VehicleDetails, error) {
	var response types.WgResponse[map[string]types.VehicleDetails]
	options := types.GetOptions(opts...)
	options.Fields = append(options.Fields, "tank_id", "name", "nation", "tier", "type", "is_premium")
	query := options.Query()

	_, err := c.Request(ctx, realm, fmt.Sprintf("encyclopedia/vehicles/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
