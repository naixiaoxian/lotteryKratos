package req

import "lotteryKratos/internal/data/activity/aggregates"

type ActivityConfigReq struct {
	ActivityId         int64
	ActivityConfigRich aggregates.ActivityConfigRich
}
