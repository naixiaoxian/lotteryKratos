package vo

import "time"

type DrawAwardVO struct {
	uId          string
	awardId      string
	awardType    int //1:文字描述、2:兑换码、3:优惠券、4:实物奖品
	awardName    string
	awardContent string
	strategyMode int
	grantType    int
	grantDate    time.Time
}

func NewDrawAwardVo(uid string, awardId string, awardType int, awardName string, awardContent string, strategyMode int, grantType int, grantDate time.Time) *DrawAwardVO {
	return &DrawAwardVO{
		uId:          uid,
		awardId:      awardId,
		awardType:    awardType,
		awardName:    awardName,
		awardContent: awardContent,
		strategyMode: strategyMode,
		grantType:    grantType,
		grantDate:    grantDate,
	}
}
