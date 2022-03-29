package algorithm

import (
	"lotteryKratos/internal/biz"
	"lotteryKratos/internal/data/strategy/vo"
)

type EntiretyRateRandomDrawAlgorithm struct {
	BaseAlgorithm BaseAlgorithm
}

func (ba *EntiretyRateRandomDrawAlgorithm) RandomDraw(strategyId int64, excludeAwardIds []string) (ret string) {
	differenceDenominator := float32(0)
	differentAwardRateList := make([]vo.AwardRateVo, 0)
	awardRateIntervalValList := ba.BaseAlgorithm.AwardRateInfoMap[strategyId]
	// 排除掉不在抽奖范围的奖品ID集合
	for _, rateVo := range awardRateIntervalValList {
		awardId := rateVo.AwardId
		if IsExistInArray(awardId, excludeAwardIds) {
			continue
		}
		differentAwardRateList = append(differentAwardRateList, rateVo)
		differenceDenominator += rateVo.GetAwardRate()
	}
	//前置判断:奖品列表为0，返回null

	if len(differentAwardRateList) == 0 {
		return "未中奖"
	}

	if len(differentAwardRateList) == 1 {
		return differentAwardRateList[0].AwardId
	}

	randomVal := biz.RInt(100)

	var awardId string
	cursorVal := 0

	for _, rateVo := range differentAwardRateList {
		rateVal := int(rateVo.GetAwardRate() / differenceDenominator * 100)
		if randomVal <= (cursorVal + rateVal) {
			awardId = rateVo.AwardId
			break
		}
		cursorVal += rateVal
	}
	return awardId

}

func (ba *EntiretyRateRandomDrawAlgorithm) IsExist(strategyId int64) bool {
	return ba.BaseAlgorithm.isExist(int64(strategyId))
}

func (ba *EntiretyRateRandomDrawAlgorithm) InitRateTuple(strategyId int64, strategyMode int, awardRateInfoList []vo.AwardRateVo) {
	ba.BaseAlgorithm.initRateTuple(strategyId, strategyMode, awardRateInfoList)
}
