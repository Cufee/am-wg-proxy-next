package shims

import "github.com/byvko-dev/am-core/helpers/env"

var legacyApiKey string
var usersApiUrl string

func init() {
	usersApiUrl = env.MustGet("LEGACY_USERS_API_URL")[0].(string)
	legacyApiKey = env.MustGet("LEGACY_API_KEY")[0].(string)
}
