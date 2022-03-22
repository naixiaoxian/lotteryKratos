package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"lotteryKratos/internal/biz"
	"lotteryKratos/internal/data/gormModel"
)

type StrategyRepoImpl interface {
	QueryExcludeAwardIds(id int64) []gormModel.StrategyDetail
	DeductStock(strategyId int64, awardId int64) error
	QueryAwardInfoByAwardId(id int64) (gormModel.Award, error)
}

type strategyRepo struct {
	data *Data
	log  *log.Helper
}

func NewStrategyRepo(data *Data, logger log.Logger) biz.ActivityRepoImpl {
	return &activityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (sr *strategyRepo) QueryExcludeAwardIds(id int64) (strategyDetails []gormModel.StrategyDetail) {
	sr.data.Db.Where("strategyId = ? ", id).Where("awardCount = ? ", 0).Find(&strategyDetails)
	return
}

func (sr *strategyRepo) DeductStock(strategyId int64, awardId int64) (err error) {
	var strategyDetail gormModel.StrategyDetail
	err = sr.data.Db.Where("strategyId = ?", strategyId).Where("awardId = ? ", awardId).First(&strategyDetail).Error
	if err != nil {
		return
	}
	//高并发情况下会出现问题。需要dbstatement
	//
	err = sr.data.Db.Exec("UPDATE strategy_detail SET award_count = award_count - 1 WHERE id = ?", strategyDetail.Id).Error
	return
}
func (sr *strategyRepo) QueryAwardInfoByAwardId(awardId int) (ret gormModel.Award, err error) {
	err = sr.data.Db.Where("id = ?", awardId).First(ret).Error
	if err != nil {
		return
	}
	return
}
