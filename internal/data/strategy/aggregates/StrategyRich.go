package aggregates

import "lotteryKratos/internal/data/strategy/vo"

type StrategyRich struct {
	StrategyId      int64
	StrategyBriefVo vo.StrategyBriefVO
	StrategyDetails []vo.StrategyDetailBriefVO
}
