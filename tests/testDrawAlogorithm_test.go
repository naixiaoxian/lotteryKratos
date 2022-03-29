package tests

import (
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"lotteryKratos/internal/biz/algorithm"
	"lotteryKratos/internal/biz/draw"
	"lotteryKratos/internal/conf"
	"lotteryKratos/internal/data"
	"lotteryKratos/internal/data/strategy/req"
	"lotteryKratos/internal/data/strategy/vo"
	"os"
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
	entire := algorithm.EntiretyRateRandomDrawAlgorithm{
		BaseAlgorithm: bareps,
	}
	entire.InitRateTuple(100001, 1, voList)

	awardIds := []string{
		"二等奖: iphone",
		"一等奖: IMac",
	}
	for i := 0; i < 20; i++ {
		fmt.Println("中奖结果 ：", entire.RandomDraw(100001, awardIds))
	}
	rateTupleMap2 := make(map[int64][]string)
	awardRateInfoMap2 := make(map[int64][]vo.AwardRateVo)
	bareps2 := algorithm.BaseAlgorithm{
		RateTupleMap:     rateTupleMap2,
		AwardRateInfoMap: awardRateInfoMap2,
	}
	singleDraw := algorithm.SingleRateRandomDrawAlgorithm{
		BaseAlgorithm: bareps2,
	}
	awardId2s := []string{
		"二等奖: iphone",
		"四等奖: Airpods",
	}
	singleDraw.InitRateTuple(100002, 2, voList)
	for i := 0; i < 20; i++ {
		fmt.Println("中奖结果 ：", singleDraw.RandomDraw(100002, awardId2s))
	}

}

func TestDraw2(t *testing.T) {
	voList := make([]vo.AwardRateVo, 0)
	voList = append(voList, vo.AwardRateVo{
		AwardId:   "1",
		AwardRate: 0.05,
	})
	voList = append(voList, vo.AwardRateVo{
		AwardId:   "2",
		AwardRate: 0.15,
	})
	voList = append(voList, vo.AwardRateVo{
		AwardId:   "3",
		AwardRate: 0.2,
	})
	voList = append(voList, vo.AwardRateVo{
		AwardId:   "4",
		AwardRate: 0.25,
	})
	voList = append(voList, vo.AwardRateVo{
		AwardId:   "5",
		AwardRate: 0.35,
	})
	rateTupleMap := make(map[int64][]string)
	awardRateInfoMap := make(map[int64][]vo.AwardRateVo)

	bareps := algorithm.BaseAlgorithm{
		RateTupleMap:     rateTupleMap,
		AwardRateInfoMap: awardRateInfoMap,
	}
	entire := algorithm.EntiretyRateRandomDrawAlgorithm{
		BaseAlgorithm: bareps,
	}
	entire.InitRateTuple(10001, 1, voList)

	rateTupleMap2 := make(map[int64][]string)
	awardRateInfoMap2 := make(map[int64][]vo.AwardRateVo)
	bareps2 := algorithm.BaseAlgorithm{
		RateTupleMap:     rateTupleMap2,
		AwardRateInfoMap: awardRateInfoMap2,
	}
	singleDraw := algorithm.SingleRateRandomDrawAlgorithm{
		BaseAlgorithm: bareps2,
	}
	singleDraw.InitRateTuple(10002, 2, voList)

	//init
	var flagconf string
	flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	dataData, _, _ := data.NewData(bc.Data, logger)
	strategyImpl := data.NewStrategyRepo(dataData, logger)

	data := draw.NewDraBaseDomain(strategyImpl, logger, singleDraw, entire)
	data.DoDrawExec(req.DrawReq{
		UId:        "kd",
		StrategyId: 10001,
		Uuid:       "kd1",
	})
	data.DoDrawExec(req.DrawReq{
		UId:        "kd2",
		StrategyId: 10001,
		Uuid:       "kd1",
	})
	data.DoDrawExec(req.DrawReq{
		UId:        "kd3",
		StrategyId: 10001,
		Uuid:       "kd1",
	})
	data.DoDrawExec(req.DrawReq{
		UId:        "kd4",
		StrategyId: 10001,
		Uuid:       "kd1",
	})
	data.DoDrawExec(req.DrawReq{
		UId:        "kd5",
		StrategyId: 10001,
		Uuid:       "kd1",
	})
}
