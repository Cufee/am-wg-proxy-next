package client

import (
	"context"
	"strconv"
	"time"

	"github.com/cufee/am-wg-proxy-next/v2/internal/api"
	"github.com/cufee/am-wg-proxy-next/v2/internal/client"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type baseClient interface {
	SearchAccounts(ctx context.Context, realm, query string, fields ...string) ([]types.Account, error)
	AccountByID(ctx context.Context, realm string, id string, fields ...string) (types.ExtendedAccount, error)
	BatchAccountByID(ctx context.Context, realm string, ids []string, fields ...string) (map[string]types.ExtendedAccount, error)
	AccountClan(ctx context.Context, realm string, id string, fields ...string) (types.ClanMember, error)
	BatchAccountClan(ctx context.Context, realm string, ids []string, fields ...string) (map[string]types.ClanMember, error)
	AccountVehicles(ctx context.Context, realm string, id string, fields ...string) ([]types.VehicleStatsFrame, error)
	AccountAchievements(ctx context.Context, realm string, id string, fields ...string) (types.AchievementsFrame, error)
	BatchAccountAchievements(ctx context.Context, realm string, ids []string, fields ...string) (map[string]types.AchievementsFrame, error)

	SearchClans(ctx context.Context, realm, query string, fields ...string) ([]types.Clan, error)
	ClanByID(ctx context.Context, realm string, id string, fields ...string) (types.ExtendedClan, error)
	BatchClanByID(ctx context.Context, realm string, ids []string, fields ...string) (map[string]types.ExtendedClan, error)

	VehicleGlossary(ctx context.Context, realm string, vehicleId string, lang string, fields ...string) (types.VehicleDetails, error)
	CompleteVehicleGlossary(ctx context.Context, realm string, lang string, fields ...string) (map[string]types.VehicleDetails, error)
}

type Client interface {
	baseClient
	RealmFromAccountID(id string) string
}

type clientWithCommon struct {
	baseClient
}

func (c clientWithCommon) RealmFromAccountID(id string) string {
	intID, _ := strconv.Atoi(id)
	switch {
	case intID == 0:
		return ""
	case intID < 500000000:
		return "RU"
	case intID < 1000000000:
		return "EU"
	case intID < 2000000000:
		return "NA"
	default:
		return "AS"
	}
}

func NewEmbeddedClient(primaryWgAppID string, primaryWgAppRPS int, proxyHostList string, requestTimeout time.Duration) (Client, error) {
	c, err := client.NewClient(primaryWgAppID, primaryWgAppRPS, client.Options{BucketsString: proxyHostList, Timeout: requestTimeout})
	if err != nil {
		return nil, err
	}
	return clientWithCommon{c}, nil
}

func NewRemoteClient(apiHost string, requestTimeout time.Duration) (Client, error) {
	return clientWithCommon{api.NewClient(apiHost, requestTimeout)}, nil
}
