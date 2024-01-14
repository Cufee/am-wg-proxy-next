package glossary

import (
	"errors"
	"fmt"

	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
	"github.com/cufee/am-wg-proxy-next/types"
)

type GlossaryVehicleResponse struct {
	types.WgResponse
	Data map[string]types.VehicleDetails `json:"data"`
}

func GetGlossaryVehicle(realm string, vehicleID, lang string) (*types.VehicleDetails, error) {
	var response GlossaryVehicleResponse
	_, err := client.WargamingRequest(realm, fmt.Sprintf("encyclopedia/vehicles/?tank_id=%v&fields=tank_id,name,nation,tier,type,is_premium&language=%v", vehicleID, lang), "GET", nil, &response)
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

func GetAllGlossaryVehicles(realm, lang string) (map[string]types.VehicleDetails, error) {
	var response GlossaryVehicleResponse
	_, err := client.WargamingRequest(realm, fmt.Sprintf("encyclopedia/vehicles/?fields=tank_id,name,nation,tier,type,is_premium&language=%v", lang), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
