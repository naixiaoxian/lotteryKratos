package algorithm

import (
	"lotteryKratos/internal/biz"
	"lotteryKratos/internal/data/strategy/vo"
)

type SingleRateRandomDrawAlgorithm struct {
	BaseAlgorithm BaseAlgorithm
}

func (ba *SingleRateRandomDrawAlgorithm) RandomDraw(strategyId int64, excludeAwardIds []string) (ret string) {
	rateTuple, existed := ba.BaseAlgorithm.RateTupleMap[strategyId]
	if !existed {
		return ""
	}
	randomVal := biz.RInt(100)
	idx := ba.BaseAlgorithm.hashIdx(randomVal)
	ret = rateTuple[idx]

	if IsExistInArray(ret, excludeAwardIds) {
		ret = "未中奖"
	}
	return
}

func (ba *SingleRateRandomDrawAlgorithm) IsExist(strategyId int64) bool {
	return ba.BaseAlgorithm.isExist(strategyId)
}

func (ba *SingleRateRandomDrawAlgorithm) InitRateTuple(strategyId int64, strategyMode int, awardRateInfoList []vo.AwardRateVo) {
	ba.BaseAlgorithm.initRateTuple(strategyId, strategyMode, awardRateInfoList)
}
