package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"lotteryKratos/internal/biz/strategy/draw"
	"lotteryKratos/internal/data/gormModel"
	"lotteryKratos/internal/data/strategy/aggregates"
	"lotteryKratos/internal/data/strategy/vo"
)

type strategyRepo struct {
	data *Data
	log  *log.Helper
}

func NewStrategyRepo(data *Data, logger log.Logger) draw.StrategyRepoImpl {
	return &strategyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (sr *strategyRepo) QueryExcludeAwardIds(id int64) (strategyDetails []gormModel.StrategyDetail) {
	sr.data.Db.Where("strategy_id = ? ", id).Where("award_surplus_count = ? ", 0).Find(&strategyDetails)
	return
}

func (sr *strategyRepo) DeductStock(strategyId int64, awardId int64) (err error) {
	var strategyDetail gormModel.StrategyDetail
	err = sr.data.Db.Where("strategy_id = ?", strategyId).Where("award_id = ? ", awardId).First(&strategyDetail).Error
	if err != nil {
		return
	}
	//高并发情况下会出现问题。需要dbstatement
	err = sr.data.Db.Exec("UPDATE strategy_detail SET award_surplus_count = award_surplus_count - 1 WHERE id = ?", strategyDetail.Id).Error
	return
}

func (sr *strategyRepo) QueryAwardInfoByAwardId(awardId int64) (ret gormModel.Award, err error) {
	err = sr.data.Db.Where("id = ? ", awardId).Find(&ret).Error
	if err != nil {
		return
	}
	return
}

func (sr *strategyRepo) QueryStrategyRich(strategyId int64) (ret aggregates.StrategyRich, err error) {
	var sModel gormModel.Strategy
	var details []gormModel.StrategyDetail
	err = sr.data.Db.Where("strategy_id = ?", strategyId).First(&sModel).Error
	if err != nil {
		return
	}
	err = sr.data.Db.Where("strategy_id = ?", strategyId).Find(&details).Error
	if err != nil {
		return
	}
	briefVo := vo.StrategyBriefVO{
		StrategyId:   sModel.StrategyId,
		StrategyDesc: sModel.StrategyDesc,
		StrategyMode: sModel.StrategyMode,
		GrantType:    sModel.GrantType,
		GrandDate:    sModel.GrantDate,
		ExtInfo:      sModel.ExtraInfo,
	}
	detailVoList := make([]vo.StrategyDetailBriefVO, 0)
	for _, strategyDetail := range details {
		rate, _ := strategyDetail.AwardRate.Float64()
		detailVo := vo.StrategyDetailBriefVO{
			StrategyId:        strategyDetail.StrategyId,
			AwardId:           string(rune(strategyDetail.AwardId)),
			AwardName:         strategyDetail.AwardDesc,
			AwardCount:        strategyDetail.AwardCount,
			AwardSurplusCount: strategyDetail.AwardCount,
			AwardRate:         float32(rate),
		}
		detailVoList = append(detailVoList, detailVo)
	}
	ret.StrategyId = strategyId
	ret.StrategyBriefVo = briefVo
	ret.StrategyDetails = detailVoList
	return
}
