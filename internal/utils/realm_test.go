package utils

import (
	"testing"

	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/matryer/is"
)

func TestRealmFromID(t *testing.T) {
	is := is.New(t)
	t.Run("bot", func(t *testing.T) {
		realm := RealmFromID("4200000001")
		is.True(*realm == types.RealmBotAccounts)
	})
	t.Run("china", func(t *testing.T) {
		realm := RealmFromID("3100000001")
		is.True(*realm == types.RealmChina)
	})
	t.Run("asia", func(t *testing.T) {
		realm := RealmFromID("2000000001")
		is.True(*realm == types.RealmAsia)
	})
	t.Run("na", func(t *testing.T) {
		realm := RealmFromID("1000000001")
		is.True(*realm == types.RealmNorthAmerica)
	})
	t.Run("eu", func(t *testing.T) {
		realm := RealmFromID("500000001")
		is.True(*realm == types.RealmEurope)
	})
	t.Run("ru", func(t *testing.T) {
		realm := RealmFromID("100000001")
		is.True(*realm == types.RealmRussia)
	})
}
