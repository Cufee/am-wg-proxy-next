package accounts

type TokenRefreshResponse struct {
	AccessToken string `json:"access_token"`
	AccountID   int    `json:"account_id"`
	Expiration  int    `json:"expires_at"`
}
