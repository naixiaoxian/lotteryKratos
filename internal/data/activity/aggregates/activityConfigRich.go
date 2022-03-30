package aggregates

import "lotteryKratos/internal/data/activity/vo"

type ActivityConfigRich struct {
	Activity  vo.ActivityVO
	Strategy  vo.StrategyVO
	AwardList []vo.AwardVO
}
