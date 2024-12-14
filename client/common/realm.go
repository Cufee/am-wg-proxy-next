package common

import (
	"strconv"
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type realmParams struct {
	realm types.Realm
	min   int
}

var regions = []realmParams{
	{realm: types.RealmBotAccounts, min: 42e8},
	{realm: types.RealmChina, min: 31e8},
	{realm: types.RealmAsia, min: 20e8},
	{realm: types.RealmNorthAmerica, min: 10e8},
	{realm: types.RealmEurope, min: 5e8},
	{realm: types.RealmRussia, min: 0e8},
}

func RealmFromID(idStr string) (types.Realm, bool) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return types.Realm(""), false
	}
	for _, region := range regions {
		if id >= region.min {
			return region.realm, true
		}
	}
	return types.Realm(""), false
}

func ParseRealm(realm string) (types.Realm, bool) {
	for _, region := range regions {
		if region.realm.String() == strings.ToUpper(realm) {
			return region.realm, true
		}
	}
	return types.Realm(""), false
}
