package accounts

import "aftermath.link/repo/am-wg-proxy/handlers"

type AccountInfoResponse struct {
	handlers.WargamingBaseResponse
	Data map[string]AccountExtendedInfo `json:"data"`
}

type AccountExtendedInfo struct {
	AccountBaseInfo
	Statistics     interface{} `json:"statistics"`
	CreatedAt      int64       `json:"created_at"`
	UpdatedAt      int64       `json:"updated_at"`
	Private        interface{} `json:"private"`
	LastBattleTime int64       `json:"last_battle_time"`
}
