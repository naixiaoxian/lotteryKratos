package vo

import "time"

type drawOrderVO struct {
	UId          string
	TakeId       int64
	ActivityId   int64
	OrderId      int64
	StrategyId   int64
	StrategyMode int
	GrantType    int
	GrantDate    time.Time
	GrantState   int
	AwardId      string
	//1:文字描述、2:兑换码、3:优惠券、4:实物奖品
	AwardType    int
	AwardName    string
	AwardContent string
}
