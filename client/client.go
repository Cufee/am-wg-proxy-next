package client

import (
	"context"
	"os"
	"time"

	"github.com/cufee/am-wg-proxy-next/v2/internal/api"
	"github.com/cufee/am-wg-proxy-next/v2/internal/client"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/rs/zerolog"
)

type baseClient interface {
	SearchAccounts(ctx context.Context, realm types.Realm, query string, opts ...types.Option) ([]types.Account, error)
	AccountByID(ctx context.Context, realm types.Realm, id string, opts ...types.Option) (types.ExtendedAccount, error)
	BatchAccountByID(ctx context.Context, realm types.Realm, ids []string, opts ...types.Option) (map[string]types.ExtendedAccount, error)
	AccountClan(ctx context.Context, realm types.Realm, id string, opts ...types.Option) (types.ClanMember, error)
	BatchAccountClan(ctx context.Context, realm types.Realm, ids []string, opts ...types.Option) (map[string]types.ClanMember, error)
	AccountVehicles(ctx context.Context, realm types.Realm, id string, vehicles []string, opts ...types.Option) ([]types.VehicleStatsFrame, error)
	AccountAchievements(ctx context.Context, realm types.Realm, id string, opts ...types.Option) (types.AchievementsFrame, error)
	AccountVehicleAchievements(ctx context.Context, realm types.Realm, id string, opts ...types.Option) (map[string]types.AchievementsFrame, error)
	BatchAccountAchievements(ctx context.Context, realm types.Realm, ids []string, opts ...types.Option) (map[string]types.AchievementsFrame, error)

	SearchClans(ctx context.Context, realm types.Realm, query string, opts ...types.Option) ([]types.Clan, error)
	ClanByID(ctx context.Context, realm types.Realm, id string, opts ...types.Option) (types.ExtendedClan, error)
	BatchClanByID(ctx context.Context, realm types.Realm, ids []string, opts ...types.Option) (map[string]types.ExtendedClan, error)

	VehicleGlossary(ctx context.Context, realm types.Realm, vehicleId string, opts ...types.Option) (types.VehicleDetails, error)
	CompleteVehicleGlossary(ctx context.Context, realm types.Realm, opts ...types.Option) (map[string]types.VehicleDetails, error)

	ParseRealm(value string) (types.Realm, bool)
	RealmFromID(id string) (types.Realm, bool)
}

type Client interface {
	baseClient
}

type clientOptions struct {
	logLevel zerolog.Level
}

var defaultOptions = clientOptions{logLevel: zerolog.InfoLevel}

type ClientOption func(*clientOptions)

func WithLogLevel(level zerolog.Level) ClientOption {
	return func(co *clientOptions) { co.logLevel = level }
}

func NewEmbeddedClient(primaryWgAppID string, primaryWgAppRPS int, proxyHostList string, requestTimeout time.Duration, opts ...ClientOption) (Client, error) {
	options := defaultOptions
	for _, apply := range opts {
		apply(&defaultOptions)
	}
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(options.logLevel)

	c, err := client.NewClient(logger, primaryWgAppID, primaryWgAppRPS, client.Options{BucketsString: proxyHostList, Timeout: requestTimeout})
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewRemoteClient(apiHost string, requestTimeout time.Duration, opts ...ClientOption) (Client, error) {
	options := defaultOptions
	for _, apply := range opts {
		apply(&defaultOptions)
	}
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(options.logLevel)

	return api.NewClient(logger, apiHost, requestTimeout), nil
}
