package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"lotteryKratos/internal/biz/activity"
	"lotteryKratos/internal/data/activity/aggregates"
	"lotteryKratos/internal/data/activity/req"
	"lotteryKratos/internal/data/activity/vo"
	"lotteryKratos/internal/data/gormModel"
	"strconv"
	"time"
)

type DeployRep struct {
	data *Data
	log  *log.Helper
	DB   *gorm.DB
}

func NewDeployRepo(data *Data, logger log.Logger) activity.IActivityRepImpl {
	return &DeployRep{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (d *DeployRep) BeginTrans() {
	d.DB = d.data.Db.Begin()
}

func (d *DeployRep) Commit() {
	d.DB.Commit()
	d.DB = nil
}
func (d *DeployRep) Rollback() {
	d.DB.Rollback()
	d.DB = nil
}

func (d *DeployRep) AddActivity(activityVo vo.ActivityVO) {
	deployModel := &gormModel.Activity{
		ActivityName:      activityVo.ActivityName,
		ActivityDesc:      activityVo.ActivityDesc,
		ActivityID:        activityVo.ActivityId,
		BeginDateTime:     activityVo.BeginDateTime,
		EndDateTime:       activityVo.EndDateTime,
		StockCount:        activityVo.StockCount,
		TakeCount:         activityVo.TakeCount,
		StockSurplusCount: activityVo.StockSurplusCount,
		State:             activityVo.State,
		Creator:           "",
		CreateTime:        time.Now(),
		UpdateTime:        time.Now(),
	}
	err := d.DB.Create(deployModel).Error
	if err != nil {
		panic(err)
	}
	//redis 设置活动库存
}

func (d *DeployRep) AddAward(awardList []vo.AwardVO) {
	var awardModelList []gormModel.Award
	for _, awardVO := range awardList {
		awardId, _ := strconv.Atoi(awardVO.AwardId)
		award := &gormModel.Award{
			AwardId:      int64(awardId),
			AwardType:    awardVO.AwardType,
			AwardCount:   0,
			AwardName:    awardVO.AwardName,
			AwardContent: awardVO.AwardContent,
			CreateTime:   time.Now(),
			UpdateTime:   time.Now(),
		}
		awardModelList = append(awardModelList, *award)
	}
	err := d.DB.Create(&awardModelList).Error
	if err != nil {
		panic(err)
	}
}

func (d *DeployRep) AddStrategy(strategy vo.StrategyVO) {
	strategyModel := &gormModel.Strategy{
		StrategyId:   strategy.StrategyId,
		StrategyDesc: strategy.StrategyDesc,
		StrategyMode: strategy.StrategyMode,
		GrantType:    strategy.GrantType,
		GrantDate:    strategy.GrantDate,
		ExtInfo:      strategy.ExtInfo,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	err := d.DB.Create(strategyModel).Error
	if err != nil {
		panic(err)
	}
}

func (d *DeployRep) AddStrategyDetailList(strategyDetailList []vo.StrategyDetailVO) {
	var detailModelList []gormModel.StrategyDetail
	for _, strategyDetail := range strategyDetailList {
		detail := &gormModel.StrategyDetail{
			StrategyId:        strategyDetail.StrategyId,
			AwardId:           strategyDetail.AwardId,
			AwardCount:        strategyDetail.AwardCount,
			AwardRate:         strategyDetail.AwardRate,
			AwardSurplusCount: strategyDetail.AwardSurplusCount,
			AwardName:         strategyDetail.AwardName,
			CreateTime:        time.Now(),
			UpdateTime:        time.Now(),
		}
		detailModelList = append(detailModelList, *detail)
	}
	err := d.DB.Create(detailModelList).Error
	if err != nil {
		panic(err)
	}
}

func (d *DeployRep) AlterStatus(activityID int64, beforeState int, afterState int) (ret bool) {
	var activitya gormModel.Activity
	d.data.Db.Where("id = ? ", activityID).Where("state = ? ", beforeState).First(&activitya)
	activitya.State = afterState
	count := d.data.Db.Save(activitya).RowsAffected
	ret = 1 == count
	return
}

func (d *DeployRep) QueryActivityBill(req req.PartakeReq) (ret vo.ActivityBillVO) {
	return
}

func (d *DeployRep) SubtractionActivityStock(activityId int64) (ret int) {
	return
}

func (d *DeployRep) ScanToDoActivityList(id int64) (ret []vo.ActivityVO) {
	return
}

func (d *DeployRep) SubtractionActivityStockByRedis(uid string, activityId int64, stockCount int) {
	return
}

func (d DeployRep) RecoverActivityCacheStockByRedis(activityId int64, tokenKey string, code string) {
	return
}

func (d DeployRep) QueryActivityInfoLimitPage(req req.ActivityINfoLimitPageReq) (ret aggregates.ActivityInfoLimitPageRich) {
	return
}
