package tests

import (
	"fmt"
	"lotteryKratos/internal/biz/algorithm"
	"lotteryKratos/internal/data/strategy/vo"
	"testing"
)

func TestDraw(t *testing.T) {

	voList := make([]vo.AwardRateVo, 0)
	voList = append(voList, vo.AwardRateVo{
		AwardId:   "一等奖: IMac",
		AwardRate: 0.05,
	})
	voList = append(voList, vo.AwardRateVo{
		AwardId:   "二等奖: iphone",
		AwardRate: 0.15,
	})
	voList = append(voList, vo.AwardRateVo{
		AwardId:   "三等奖: ipad",
		AwardRate: 0.2,
	})
	voList = append(voList, vo.AwardRateVo{
		AwardId:   "四等奖: Airpods",
		AwardRate: 0.25,
	})
	voList = append(voList, vo.AwardRateVo{
		AwardId:   "五等奖: 充电宝",
		AwardRate: 0.35,
	})
	rateTupleMap := make(map[int64][]string)
	awardRateInfoMap := make(map[int64][]vo.AwardRateVo)

	bareps := algorithm.BaseAlgorithm{
		RateTupleMap:     rateTupleMap,
		AwardRateInfoMap: awardRateInfoMap,
	}
	single := algorithm.EntiretyRateRandomDrawAlgorithm{
		BaseAlgorithm: bareps,
	}
	single.InitRateTuple(100001, 1, voList)

	awardIds := []string{
		"二等奖: iphone",
		"四等奖: Airpods",
	}
	for i := 0; i < 20; i++ {
		fmt.Println("中奖结果 ：", single.RandomDraw(100001, awardIds))
	}
}
