package client

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/cufee/am-wg-proxy-next/types"
)

type glossaryVehicleResponse struct {
	types.WgResponse
	Data map[string]types.VehicleDetails `json:"data"`
}

func (c *Client) GetGlossaryVehicle(realm string, vehicleID, lang string, fields ...string) (*types.VehicleDetails, error) {
	var response glossaryVehicleResponse
	query := url.Values{}
	query.Set("fields", "tank_id,name,nation,tier,type,is_premium")
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("tank_id", vehicleID)
	query.Set("language", lang)

	_, err := c.Request(realm, fmt.Sprintf("encyclopedia/vehicles/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}

	info, ok := response.Data[fmt.Sprint(vehicleID)]
	if !ok || info.TankID == 0 {
		return nil, errors.New("vehicle not found")
	}
	return &info, nil
}

func (c *Client) GetAllGlossaryVehicles(realm, lang string, fields ...string) (map[string]types.VehicleDetails, error) {
	var response glossaryVehicleResponse
	query := url.Values{}
	query.Set("fields", "tank_id,name,nation,tier,type,is_premium")
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("language", lang)

	_, err := c.Request(realm, fmt.Sprintf("encyclopedia/vehicles/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
