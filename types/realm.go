package types

import (
	"strings"
)

func ParseRealm(value string) (Realm, bool) {
	value = strings.ToUpper(value)
	switch value {
	default:
		return Realm(""), false

	case RealmBotAccounts.String():
		return RealmBotAccounts, true

	case RealmNorthAmerica.String():
		return RealmNorthAmerica, true

	case RealmEurope.String():
		return RealmEurope, true

	case RealmRussia.String():
		return RealmRussia, true

	case RealmChina.String():
		return RealmChina, true

	case RealmAsia.String():
		return RealmAsia, true
	}
}

type Realm string

func (r Realm) String() string {
	return string(r)
}

func (r Realm) DomainBlitz() (string, bool) {
	switch r {
	default:
		return "", false
	case RealmNorthAmerica:
		return "na.wotblitz.com", true
	case RealmEurope:
		return "eu.wotblitz.com", true
	case RealmAsia:
		return "asia.wotblitz.com", true
	}
}

func (r Realm) DomainWorldOfTanks() (string, bool) {
	switch r {
	default:
		return "", false
	case RealmNorthAmerica:
		return "api.worldoftanks.com", true
	case RealmEurope:
		return "api.worldoftanks.eu", true
	case RealmAsia:
		return "api.worldoftanks.asia", true
	}
}

const (
	RealmBotAccounts = Realm("BOT")

	RealmNorthAmerica = Realm("NA")
	RealmEurope       = Realm("EU")
	RealmRussia       = Realm("RU")
	RealmChina        = Realm("CH")
	RealmAsia         = Realm("AS")
)
