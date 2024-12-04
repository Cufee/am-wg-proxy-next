package types

type Realm string

func (r Realm) String() string {
	return string(r)
}

const (
	RealmBotAccounts = Realm("BOT")

	RealmNorthAmerica = Realm("NA")
	RealmEurope       = Realm("EU")
	RealmRussia       = Realm("RU")
	RealmChina        = Realm("CH")
	RealmAsia         = Realm("AS")
)
