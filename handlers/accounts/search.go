package accounts

import (
	"aftermath.link/repo/am-wg-proxy/handlers"
	"aftermath.link/repo/logs"
)

type SearchResponse struct {
	handlers.WargamingBaseResponse
	Data []AccountBaseInfo `json:"data"`
}

type AccountBaseInfo struct {
	Nickname  string `json:"nickname"`
	AccountID int    `json:"account_id"`
}

func SearchWargamingPlayers(query string) ([]AccountBaseInfo, error) {

	handlers.HTTPRequest(url, "GET", nil, nil)

	var accounts []AccountBaseInfo
	_, err := db.Search("accounts", query, &accounts)
	if err != nil {
		return nil, logs.Wrap(err, "db.Search failed")
	}

	return accounts, nil
}
