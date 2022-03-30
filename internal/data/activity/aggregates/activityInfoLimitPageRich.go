package aggregates

import "lotteryKratos/internal/data/activity/vo"

type ActivityInfoLimitPageRich struct {
	Count          int64
	ActivityVOList []vo.ActivityVO
}
