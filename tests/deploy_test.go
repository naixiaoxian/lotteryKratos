package tests

import (
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"lotteryKratos/internal/biz/activity"
	"lotteryKratos/internal/biz/activity/stateFlow"
	"lotteryKratos/internal/common"
	"lotteryKratos/internal/conf"
	"lotteryKratos/internal/data"
	"lotteryKratos/internal/data/activity/aggregates"
	"lotteryKratos/internal/data/activity/req"
	"lotteryKratos/internal/data/activity/vo"
	vo2 "lotteryKratos/internal/data/award/vo"
	"os"
	"testing"
	"time"
)

var logger log.Logger
var dataData *data.Data

func init() {
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

	logger = log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	dataData, _, _ = data.NewData(bc.Data, logger)
}

func TestDeploy_createActivity(t *testing.T) {
	activityVo := vo.ActivityVO{
		ActivityId:        120981321,
		ActivityName:      "测试活动",
		ActivityDesc:      "测试活动描述",
		BeginDateTime:     time.Now(),
		EndDateTime:       time.Now(),
		StockCount:        100,
		StockSurplusCount: 100,
		State:             1,
		StrategyId:        10002,
		TakeCount:         0,
		UserTakeLeftCount: 0,
		Creator:           "kd",
		CreateTime:        time.Now(),
		UpdateTime:        time.Now(),
	}
	strategyDetail01 := &vo.StrategyDetailVO{
		StrategyId:        10002,
		AwardId:           101,
		AwardName:         "一等奖",
		AwardCount:        10,
		AwardSurplusCount: 10,
		AwardRate:         0.05,
		ShippingAddress:   vo2.ShippingAddress{},
		ExtInfo:           "",
	}
	strategyDetail02 := &vo.StrategyDetailVO{
		StrategyId:        10002,
		AwardId:           102,
		AwardName:         "二等奖",
		AwardCount:        20,
		AwardSurplusCount: 20,
		AwardRate:         0.15,
		ShippingAddress:   vo2.ShippingAddress{},
		ExtInfo:           "",
	}
	strategyDetail03 := &vo.StrategyDetailVO{
		StrategyId:        10002,
		AwardId:           103,
		AwardName:         "三等奖",
		AwardCount:        50,
		AwardSurplusCount: 50,
		AwardRate:         0.2,
		ShippingAddress:   vo2.ShippingAddress{},
		ExtInfo:           "",
	}
	strategyDetail04 := &vo.StrategyDetailVO{
		StrategyId:        10002,
		AwardId:           104,
		AwardName:         "四等奖",
		AwardCount:        100,
		AwardSurplusCount: 100,
		AwardRate:         0.25,
		ShippingAddress:   vo2.ShippingAddress{},
		ExtInfo:           "",
	}
	strategyDetail05 := &vo.StrategyDetailVO{
		StrategyId:        10002,
		AwardId:           105,
		AwardName:         "三等奖",
		AwardCount:        500,
		AwardSurplusCount: 500,
		AwardRate:         0.35,
		ShippingAddress:   vo2.ShippingAddress{},
		ExtInfo:           "",
	}

	arr := []vo.StrategyDetailVO{
		*strategyDetail01,
		*strategyDetail02,
		*strategyDetail03,
		*strategyDetail04,
		*strategyDetail05,
	}
	strategy := &vo.StrategyVO{
		StrategyId:   10002,
		StrategyDesc: "策略描述",
		StrategyMode: 1,
		GrantType:    1,
		GrantDate:    time.Now(),
		ExtInfo:      "",
		StrategyList: arr,
	}
	award01 := &vo.AwardVO{
		AwardId:      "101",
		AwardType:    1,
		AwardName:    "电脑1",
		AwardContent: "电脑描述1",
	}
	award02 := &vo.AwardVO{
		AwardId:      "102",
		AwardType:    2,
		AwardName:    "电脑2",
		AwardContent: "电脑描述2",
	}
	award03 := &vo.AwardVO{
		AwardId:      "103",
		AwardType:    3,
		AwardName:    "电脑3",
		AwardContent: "电脑描述3",
	}
	award04 := &vo.AwardVO{
		AwardId:      "104",
		AwardType:    4,
		AwardName:    "电脑4",
		AwardContent: "电脑描述4",
	}
	award05 := &vo.AwardVO{
		AwardId:      "105",
		AwardType:    5,
		AwardName:    "电脑5",
		AwardContent: "电脑描述5",
	}
	arrList := []vo.AwardVO{
		*award01,
		*award02,
		*award03,
		*award04,
		*award05,
	}
	rich := aggregates.ActivityConfigRich{
		Activity:  activityVo,
		Strategy:  *strategy,
		AwardList: arrList,
	}
	//init
	deployRep := data.NewDeployRepo(dataData, logger)
	deplo := activity.NewDeploy(deployRep)
	req := &req.ActivityConfigReq{
		ActivityId:         10002,
		ActivityConfigRich: rich,
	}
	deplo.CreateActivity(*req)
}

func TestStateFlow(t *testing.T) {
	deployRep := data.NewDeployRepo(dataData, logger)
	stateFlow.InitGroups(deployRep)
	fmt.Println(stateFlow.GetHandler(common.ACTIVITY_STATE_EDIT).Arraignment(120981321, common.ACTIVITY_STATE_EDIT))
	fmt.Println(stateFlow.GetHandler(common.ACTIVITY_STATE_ARRAIGNMENT).CheckPass(120981321, common.ACTIVITY_STATE_ARRAIGNMENT))
	fmt.Println(stateFlow.GetHandler(common.ACTIVITY_STATE_PASS).Doing(120981321, common.ACTIVITY_STATE_PASS))
	fmt.Println(stateFlow.GetHandler(common.ACTIVITY_STATE_EDIT).CheckPass(120981321, common.ACTIVITY_STATE_EDIT))
}
