package vo

import (
	"github.com/shopspring/decimal"
	"lotteryKratos/internal/data/award/vo"
)

type StrategyDetailVO struct {
	StrategyId        string
	AwardId           int64
	AwardName         string
	AwardCount        int
	AwardSurplusCount string
	AwardRate         decimal.Decimal
	ShippingAddress   vo.ShippingAddress
	ExtInfo           string
}
