package shims

import (
	"fmt"
	"net/url"

	"github.com/byvko-dev/am-core/helpers/requests"
	"github.com/byvko-dev/am-types/users/v1"
)

type usersApiResponse struct {
	users.User
	Error string `json:"error"`
}

// CheckUserByName - Check user profile by player nickname
func CheckUserByName(name, realm string) (userData usersApiResponse, err error) {
	// Make URL
	requestURL, err := url.Parse(fmt.Sprintf("%s/players/name/%s", usersApiUrl, name))
	if err != nil {
		return userData, fmt.Errorf("users api error: %s", err.Error())
	}

	// Make headers
	headers := make(map[string]string)
	headers["x-api-key"] = legacyApiKey

	// Send request
	_, err = requests.Send(requestURL.String(), "GET", headers, nil, &userData)
	if err != nil {
		return userData, fmt.Errorf("users api error: %s", err.Error())
	}

	// Check for returned error
	if userData.Error != "" {
		err = fmt.Errorf("users api error: %s", userData.Error)
	}
	return userData, err
}
