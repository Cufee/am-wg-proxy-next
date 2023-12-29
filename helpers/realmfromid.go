package helpers

func RealmFromID(pidInt int) string {
	switch {
	case pidInt == 0:
		return ""
	case pidInt < 500000000:
		return "RU"
	case pidInt < 1000000000:
		return "EU"
	case pidInt < 2000000000:
		return "NA"
	default:
		return "AS"
	}
}
