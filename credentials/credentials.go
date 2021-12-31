package credentials

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/database"
	"aftermath.link/repo/logs"
	"github.com/google/uuid"
)

const databaseName = "credentials"
const collectionName = "credentials"
const tokenKeyFormat = "%s:%s" // realm:application_id

type TokenConfig struct {
	ApplicationID string `json:"application_id"`
	RequestLimit  int    `json:"request_limit"`
	ProxyURL      string `json:"proxy_url"`
	Realm         string `json:"realm"`
}

func (t *TokenConfig) Validate() error {
	if t.ApplicationID == "" {
		return fmt.Errorf("application_id is required")
	}
	if t.RequestLimit == 0 {
		return fmt.Errorf("request_limit is required")
	}
	if t.ProxyURL == "" {
		return fmt.Errorf("proxy_url is required")
	}
	if t.Realm == "" {
		return fmt.Errorf("realm is required")
	}
	return nil
}

type Token struct {
	PublicKey string      `json:"public_key"`
	Config    TokenConfig `json:"config"`
	Valid     bool        `json:"valid"`
}

func GenerateNewCredentialsToken(token TokenConfig) (string, error) {
	var db database.Client
	db.Open(databaseName)

	var tokenDetails Token
	_, err := db.Search(collectionName, getTokenKey(token), &tokenDetails)
	if err != nil {
		return "", logs.Wrap(err, "db.Search failed")
	}
	if tokenDetails != (Token{}) {
		return "", fmt.Errorf("token already exists")
	}

	tokenDetails.PublicKey = generateNewPublicKey()
	tokenDetails.Config = token
	tokenDetails.Valid = true

	payload, err := json.Marshal(tokenDetails)
	if err != nil {
		return "", logs.Wrap(err, "json.Marshal failed")
	}
	_, err = db.Create(collectionName, getTokenKey(tokenDetails.Config), payload, 0)
	if err != nil {
		return "", logs.Wrap(err, "db.Create failed")
	}

	return tokenDetails.PublicKey, nil
}

func GetCredentialsTokenDetails(token string) (*Token, error) {
	var db database.Client
	db.Open(databaseName)

	var tokenDetailsResult []Token
	query := make(map[string]interface{})
	query["public_key"] = token
	_, err := db.SearchValues(collectionName, query, 0, &tokenDetailsResult)
	if err != nil {
		return nil, logs.Wrap(err, "db.Read failed")
	}

	if len(tokenDetailsResult) == 0 {
		return nil, fmt.Errorf("token not found")
	}
	if len(tokenDetailsResult) > 1 {
		return nil, fmt.Errorf("multiple tokens found")
	}

	return &tokenDetailsResult[0], nil
}

func DisableCredentialsToken(token string) error {
	var db database.Client
	db.Open(databaseName)

	tokenDetails, err := GetCredentialsTokenDetails(token)
	if err != nil {
		return logs.Wrap(err, "GetCredentialsTokenDetails failed")
	}
	tokenDetails.Valid = false
	payload, err := json.Marshal(tokenDetails)
	if err != nil {
		return logs.Wrap(err, "json.Marshal failed")
	}
	_, err = db.Update(collectionName, token, payload, 0, false)
	if err != nil {
		return logs.Wrap(err, "db.Update failed")
	}

	return nil
}

func generateNewPublicKey() string {
	id := uuid.NewString()
	return base64.StdEncoding.EncodeToString([]byte(id))
}

func getTokenKey(t TokenConfig) string {
	return fmt.Sprintf(tokenKeyFormat, t.Realm, t.ApplicationID)
}
