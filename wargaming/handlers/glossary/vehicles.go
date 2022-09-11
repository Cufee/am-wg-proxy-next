package glossary

import (
	"errors"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/glossary"
)

type GlossaryVehicleResponse struct {
	api.Response
	Data map[string]glossary.VehicleDetails `json:"data"`
}

func GetGlossaryVehicle(realm string, vehicleID string) (glossary.VehicleDetails, error) {
	var response GlossaryVehicleResponse
	_, err := client.WargamingRequest(realm, fmt.Sprintf("encyclopedia/vehicles/?tank_id=%v&fields=tank_id,name,nation,tier,type,is_premium", vehicleID), "GET", nil, &response)
	if err != nil {
		return glossary.VehicleDetails{}, err
	}
	if response.Error.Code != 0 {
		return glossary.VehicleDetails{}, errors.New(response.Error.Message)
	}

	info, ok := response.Data[fmt.Sprint(vehicleID)]
	if !ok {
		return info, errors.New("vehicle not found")
	}
	return info, nil
}

func GetAllGlossaryVehicles(realm string) (map[string]glossary.VehicleDetails, error) {
	var response GlossaryVehicleResponse
	_, err := client.WargamingRequest(realm, "encyclopedia/vehicles/?fields=tank_id,name,nation,tier,type,is_premium", "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
