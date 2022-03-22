package res

import "lotteryKratos/internal/data/strategy/vo"

type DrawResult struct {
	Uid           string
	StrategyId    int64
	DrawResult    int
	DrawAwardInfo vo.DrawAwardVO
}
