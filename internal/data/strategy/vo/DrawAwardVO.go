package vo

import "time"

type DrawAwardVO struct {
	uId          string
	AwardId      string
	AwardType    int //1:文字描述、2:兑换码、3:优惠券、4:实物奖品
	AwardName    string
	AwardContent string
	StrategyMode int
	GrantType    int
	GrantDate    time.Time
}

func NewDrawAwardVo(uid string, awardId string, awardType int, awardName string, awardContent string, strategyMode int, grantType int, grantDate time.Time) *DrawAwardVO {
	return &DrawAwardVO{
		uId:          uid,
		AwardId:      awardId,
		AwardType:    awardType,
		AwardName:    awardName,
		AwardContent: awardContent,
		StrategyMode: strategyMode,
		GrantType:    grantType,
		GrantDate:    grantDate,
	}
}
