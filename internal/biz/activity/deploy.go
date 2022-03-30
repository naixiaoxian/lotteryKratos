package activity

import (
	"fmt"
	"lotteryKratos/internal/data/activity/aggregates"
	"lotteryKratos/internal/data/activity/req"
	"lotteryKratos/internal/data/activity/vo"
)

//创建活动+策略

type IActivityDeploy interface {
	CreateActivity(req req.ActivityConfigReq)
	UpdateActivity(req req.ActivityConfigReq)
	ScanToDoActivityList(id int64) []vo.ActivityVO
	QueryActivityInfoLimitPaginage(req req.ActivityINfoLimitPageReq) aggregates.ActivityInfoLimitPageRich
}

type IActivityRepImpl interface {
	BeginTrans() //用于开启事务
	Commit()
	Rollback() //事务相关在这里实现
	AddActivity(activityVo vo.ActivityVO)
	AddAward(awardList []vo.AwardVO)
	AddStrategy(strategy vo.StrategyVO)
	AddStrategyDetailList(strategyDetailList []vo.StrategyDetailVO)
	AlterStatus(activityID int64, beforeState int, afterState int) bool
	QueryActivityBill(req req.PartakeReq) vo.ActivityBillVO
	SubtractionActivityStock(activityId int64) int
	ScanToDoActivityList(id int64) []vo.ActivityVO
	SubtractionActivityStockByRedis(uid string, activityId int64, stockCount int)
	RecoverActivityCacheStockByRedis(activityId int64, tokenKey string, code string)
	QueryActivityInfoLimitPage(req req.ActivityINfoLimitPageReq) aggregates.ActivityInfoLimitPageRich
}

type Deploy struct {
	activityrepImpl IActivityRepImpl
}

func NewDeploy(impl IActivityRepImpl) IActivityDeploy {
	return &Deploy{activityrepImpl: impl}
}

func (dp *Deploy) CreateActivity(req req.ActivityConfigReq) {
	configRich := req.ActivityConfigRich
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("捕获异常")
			dp.activityrepImpl.Rollback()
		}
	}()

	dp.activityrepImpl.BeginTrans()
	//活动配置
	activity := configRich.Activity
	dp.activityrepImpl.AddActivity(activity)
	//奖品配置
	awardList := configRich.AwardList
	dp.activityrepImpl.AddAward(awardList)
	//策略配置
	strategy := configRich.Strategy
	dp.activityrepImpl.AddStrategy(strategy)
	//策略明细配置
	detailList := configRich.Strategy.StrategyList
	dp.activityrepImpl.AddStrategyDetailList(detailList)
	//
	dp.activityrepImpl.Commit()
	fmt.Println("创建活动配置完成", req.ActivityId)
	return

}

func (dp *Deploy) UpdateActivity(req req.ActivityConfigReq) {
	return
}
func (dp *Deploy) ScanToDoActivityList(id int64) (ret []vo.ActivityVO) {
	ret = dp.activityrepImpl.ScanToDoActivityList(id)
	return
}
func (dp *Deploy) QueryActivityInfoLimitPaginage(req req.ActivityINfoLimitPageReq) (ret aggregates.ActivityInfoLimitPageRich) {
	ret = dp.activityrepImpl.QueryActivityInfoLimitPage(req)
	return
}
