package utils

import (
	"slices"
	"strconv"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

var minRegionIDs = map[types.Realm]int{
	types.RealmBotAccounts:  42e8,
	types.RealmChina:        31e8,
	types.RealmAsia:         20e8,
	types.RealmNorthAmerica: 10e8,
	types.RealmEurope:       5e8,
	types.RealmRussia:       0e8,
}

type realmParams struct {
	realm types.Realm
	min   int
}

var regions []realmParams

func init() {
	for realm, min := range minRegionIDs {
		regions = append(regions, realmParams{realm, min})
	}
	slices.SortFunc(regions, func(a, b realmParams) int {
		return b.min - a.min
	})
}

func RealmFromID(idStr string) *types.Realm {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil
	}
	for _, region := range regions {
		if id >= region.min {
			return &region.realm
		}
	}
	return nil
}
