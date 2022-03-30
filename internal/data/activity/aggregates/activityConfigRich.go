package aggregates

import "lotteryKratos/internal/data/activity/vo"

type ActivityConfigRich struct {
	activity  vo.ActivityVO
	strategy  vo.StrategyVO
	awardList []vo.AwardVO
}
