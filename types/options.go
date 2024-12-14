package types

import (
	"fmt"
	"net/url"
	"strings"
)

type Options struct {
	AccessToken string   `json:"access_token"`
	Fields      []string `json:"fields"`
	Extra       []string `json:"extra"`
	Language    string   `json:"language"`
	SearchType  string   `json:"type"`
	Limit       int      `json:"limit"`

	OverwriteApplicationID int        `json:"-"`
	custom                 url.Values `json:"-"`
}

type Option func(*Options)

func WithCustom(key string, value string) Option {
	return func(o *Options) {
		o.custom.Set(key, value)
	}
}

func WithFields(fields ...string) Option {
	return func(o *Options) {
		o.Fields = fields
	}
}

func WithExtraFields(fields ...string) Option {
	return func(o *Options) {
		o.Extra = fields
	}
}

func WithLanguage(language string) Option {
	return func(o *Options) {
		o.Language = language
	}
}

func WithType(searchType string) Option {
	return func(o *Options) {
		o.SearchType = searchType
	}
}

func WithToken(token string) Option {
	return func(o *Options) {
		o.AccessToken = token
	}
}

func WithLimit(limit int) Option {
	return func(o *Options) {
		o.Limit = limit
	}
}

func GetOptions(opts ...Option) Options {
	options := Options{Limit: 3, custom: make(url.Values)}
	for _, apply := range opts {
		apply(&options)
	}
	return options
}

func (o Options) Query() url.Values {
	form := o.custom
	if o.Fields != nil {
		form.Set("fields", strings.Join(o.Fields, ","))
	}
	if o.Extra != nil {
		form.Set("extra", strings.Join(o.Extra, ","))
	}
	if o.Language != "" {
		form.Set("language", o.Language)
	}
	if o.Limit > 0 {
		form.Set("limit", fmt.Sprint(o.Limit))
	}
	if o.AccessToken != "" {
		form.Set("access_token", o.AccessToken)
	}
	if o.SearchType != "" {
		form.Set("type", o.SearchType)
	}
	return form
}
