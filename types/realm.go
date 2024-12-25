package types

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

func (r Realm) DomainBlitzAPI() (string, bool) {
	switch r {
	default:
		return "", false
	case RealmNorthAmerica:
		return "api.wotblitz.com", true
	case RealmEurope:
		return "api.wotblitz.eu", true
	case RealmAsia:
		return "api.wotblitz.asia", true
	}
}

func (r Realm) DomainWorldOfTanksAPI() (string, bool) {
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
