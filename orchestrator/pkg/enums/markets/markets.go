package markets

import (
	"fmt"
	"time"
)

type Market string

const (
	MarketNse    Market = "NSE"
	MarketMcx    Market = "MCX"
	MarketNfo    Market = "NFO"
	MarketCrypto Market = "Crypto"
)

func (s Market) String() string {
	return string(s)
}

func (m Market) Timezone() (*time.Location, error) {
	var MarketTimezones = map[Market]string{
		MarketNse:    "Asia/Kolkata",
		MarketMcx:    "Asia/Kolkata",
		MarketNfo:    "Asia/Kolkata",
		MarketCrypto: "UTC",
	}

	tz, ok := MarketTimezones[m]
	if !ok {
		return nil, fmt.Errorf("invalid market: %s", m)
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, fmt.Errorf("failed to load location %q: %w", tz, err)
	}

	return loc, nil
}

func Parsemarket(s string) (Market, error) {
	switch s {
	case string(MarketNse):
		return MarketNse, nil
	case string(MarketMcx):
		return MarketMcx, nil
	case string(MarketNfo):
		return MarketNfo, nil
	case string(MarketCrypto):
		return MarketCrypto, nil
	default:
		return "", fmt.Errorf("invalid market: %s", s)
	}
}
