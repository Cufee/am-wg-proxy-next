package database

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"aftermath.link/repo/logs"
)

// Schema for responses from DB API
type ApiResponse struct {
	Success int         `json:"success"`
	Error   int         `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Errors
var (
	ErrAlreadyExists = errors.New("already exists")
	ErrEmptyDocument = errors.New("empty document")
	ErrNoCollection  = errors.New("no collection")
	ErrNoDatabase    = errors.New("no database")
	ErrNotFound      = errors.New("not found")

	ErrDocumentExpired = errors.New("document expired")
	ErrInvalidDocument = errors.New("invalid document")
	ErrInvalidID       = errors.New("invalid id")
)

var accessKey string
var baseApiUri string

func init() {
	accessKey = os.Getenv("DATABASE_ACCESS_KEY")
	if accessKey == "" {
		panic("DATABASE_ACCESS_KEY is not set")
	}
	baseApiUri = os.Getenv("DATABASE_API_URI")
	if baseApiUri == "" {
		panic("DATABASE_API_URI is not set")
	}
}

type Client struct {
	database string
}

func (c *Client) Open(database string) {
	c.database = database
}

type apiRequest struct {
	Query      string        `json:"query"`
	Database   string        `json:"database"`
	Collection string        `json:"collection"`
	Key        string        `json:"key"`
	Value      string        `json:"value"`
	Expiration time.Duration `json:"expiration"`
	Upsert     bool          `json:"upsert"`
	Operation  string        `json:"operation"`
	Limit      int           `json:"limit"`
}

const (
	createManyEndoint  = "/createMany"
	createEndoint      = "/create"
	readEndpoint       = "/read"
	updateEndpoint     = "/update"
	deleteEndpoint     = "/delete"
	searchEndpoint     = "/search"
	searchKeysEndpoint = "/search/keys"
)

func (c *Client) Create(collectionName string, key string, value []byte, expiration time.Duration) (*ApiResponse, error) {
	var url string = baseApiUri + createEndoint
	var request apiRequest
	request.Database = c.database
	request.Collection = collectionName
	request.Key = key
	request.Value = string(value)
	request.Expiration = expiration
	request.Operation = "create"

	var result ApiResponse
	err := httpRequest(url, "POST", request, &result)
	if err != nil {
		return nil, logs.Wrap(err, "httpRequest failed")
	}
	return &result, nil
}

func (c *Client) CreateMany(collectionName string, paylad map[string][]byte, expiration time.Duration) (*ApiResponse, error) {
	json, err := json.Marshal(paylad)
	if err != nil {
		return nil, err
	}
	var url string = baseApiUri + createManyEndoint
	var request apiRequest
	request.Database = c.database
	request.Collection = collectionName
	request.Value = string(json)
	request.Expiration = expiration
	request.Operation = "createMany"

	var result ApiResponse
	err = httpRequest(url, "POST", request, &result)
	if err != nil {
		return nil, logs.Wrap(err, "httpRequest failed")
	}
	return &result, nil
}

func (c *Client) Read(collectionName string, key string, out interface{}) (*ApiResponse, error) {
	var url string = baseApiUri + readEndpoint
	var request apiRequest
	request.Database = c.database
	request.Collection = collectionName
	request.Key = key
	request.Operation = "read"

	var result ApiResponse
	err := httpRequest(url, "POST", request, &result)
	if err != nil {
		return nil, logs.Wrap(err, "httpRequest failed")
	}
	return &result, decodeInterfaceToStruct(result, out)
}

func (c *Client) Update(collectionName string, key string, value []byte, expiration time.Duration, upsert bool) (*ApiResponse, error) {
	var url string = baseApiUri + updateEndpoint
	var request apiRequest
	request.Database = c.database
	request.Collection = collectionName
	request.Key = key
	request.Upsert = upsert
	request.Value = string(value)
	request.Expiration = expiration
	request.Operation = "update"

	var result ApiResponse
	err := httpRequest(url, "POST", request, &result)
	if err != nil {
		return nil, logs.Wrap(err, "httpRequest failed")
	}
	return &result, nil
}
func (c *Client) Delete(collectionName string, key string) (*ApiResponse, error) {
	var url string = baseApiUri + deleteEndpoint
	var request apiRequest
	request.Database = c.database
	request.Collection = collectionName
	request.Key = key
	request.Operation = "delete"

	var result ApiResponse
	err := httpRequest(url, "POST", request, &result)
	if err != nil {
		return nil, logs.Wrap(err, "httpRequest failed")
	}
	return &result, nil
}

func (c *Client) Search(collectionName string, query string, out interface{}) (*ApiResponse, error) {
	var url string = baseApiUri + searchEndpoint
	var request apiRequest
	request.Database = c.database
	request.Collection = collectionName
	request.Query = query
	request.Operation = "search"

	var result ApiResponse
	err := httpRequest(url, "POST", request, &result)
	if err != nil {
		return nil, logs.Wrap(err, "httpRequest failed")
	}
	return &result, decodeInterfaceToStruct(result, out)
}

func (c *Client) SearchKeys(collectionName string, query string, limit int, out interface{}) (*ApiResponse, error) {
	var url string = baseApiUri + searchKeysEndpoint
	var request apiRequest
	request.Database = c.database
	request.Collection = collectionName
	request.Query = query
	request.Operation = "search"
	request.Limit = limit

	var result ApiResponse
	err := httpRequest(url, "POST", request, &result)
	if err != nil {
		return nil, logs.Wrap(err, "httpRequest failed")
	}
	return &result, decodeInterfaceToStruct(result, out)
}

func (c *Client) SearchValues(collectionName string, query map[string]interface{}, limit int, out interface{}) (*ApiResponse, error) {
	queryString, err := json.Marshal(query)
	if err != nil {
		return nil, logs.Wrap(err, "failed to marshal query")
	}

	var url string = baseApiUri + searchKeysEndpoint
	var request apiRequest
	request.Database = c.database
	request.Collection = collectionName
	request.Query = string(queryString)
	request.Operation = "searchValues"
	request.Limit = limit

	var result ApiResponse
	err = httpRequest(url, "POST", request, &result)
	if err != nil {
		return nil, logs.Wrap(err, "httpRequest failed")
	}
	return &result, decodeInterfaceToStruct(result, out)
}

func decodeInterfaceToStruct(in interface{}, out interface{}) error {
	if in == nil {
		return nil
	}
	if out == nil {
		return nil
	}

	data, err := json.Marshal(in)
	if err != nil {
		return logs.Wrap(err, "failed to marshal data")
	}
	err = json.Unmarshal(data, out)
	if err != nil {
		return logs.Wrap(err, "failed to unmarshal data")
	}

	return nil
}
