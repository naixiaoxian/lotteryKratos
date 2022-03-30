package vo

import (
	"lotteryKratos/internal/data/award/vo"
)

type StrategyDetailVO struct {
	StrategyId        int64
	AwardId           int
	AwardName         string
	AwardCount        int
	AwardSurplusCount int
	AwardRate         float32
	ShippingAddress   vo.ShippingAddress
	ExtInfo           string
}
