package draw

import (
	"github.com/go-kratos/kratos/v2/log"
	"lotteryKratos/internal/biz/algorithm"
	"lotteryKratos/internal/data"
	"lotteryKratos/internal/data/gormModel"
	"lotteryKratos/internal/data/strategy/aggregates"
	"lotteryKratos/internal/data/strategy/req"
	"lotteryKratos/internal/data/strategy/res"
	"lotteryKratos/internal/data/strategy/vo"
	"strconv"
)

type IDrawExec interface {
	doDrawExec(req req.DrawReq) res.DrawResult
}

//func init() {
//	//初始化。根据algorithmGroup 单条配置
//	//algorithmGroup :=
//}

type DrawBase struct {
	rep        data.StrategyRepoImpl
	log        log.Helper
	singleDraw algorithm.SingleRateRandomDrawAlgorithm
	entireDraw algorithm.EntiretyRateRandomDrawAlgorithm
}

func (db *DrawBase) getAlgorithm(mode int) algorithm.DrawImpl {
	if mode == 1 {
		return &db.entireDraw
	}
	return &db.singleDraw
}

func (db *DrawBase) doDrawExec(req req.DrawReq) (ret *res.DrawResult) {
	stragetyRich := aggregates.StrategyRich{}
	strategyBrief := stragetyRich.StrategyBriefVo
	db.checkAndInitRateData(req.StrategyId, strategyBrief.StrategyMode, stragetyRich.StrategyDetails)
	//抽奖算法
	exludeAwardIds := db.queryExcludeAwardIds(req.StrategyId)
	//包装中奖结果
	awardId := db.drawAlgorithm(req.StrategyId, &db.entireDraw, exludeAwardIds)

	return db.buildDrawResult(req.UId, req.StrategyId, awardId, strategyBrief)
}

func (db *DrawBase) checkAndInitRateData(strategyId int64, strategyMode int, strategyDetailList []vo.StrategyDetailBriefVO) {
	//strategyRich :=
	algom := db.getAlgorithm(strategyMode)
	if algom.IsExist(strategyId) {
		return
	}
	var awardRateInfo []vo.AwardRateVo
	for _, briefVO := range strategyDetailList {
		awardRateInfo = append(awardRateInfo, vo.AwardRateVo{
			AwardId:   briefVO.AwardId,
			AwardRate: briefVO.AwardRate,
		})
	}
	algom.InitRateTuple(strategyId, strategyMode, awardRateInfo)
}

func (db *DrawBase) drawAlgorithm(strategyId int64, drawImpl algorithm.DrawImpl, strategyDetailList []gormModel.StrategyDetail) (awardId string) {
	//strategyRich :=
	ids := make([]string, 0)
	for _, strategyDetail := range strategyDetailList {
		ids = append(ids, string(strategyDetail.AwardId))
	}
	awardId = drawImpl.RandomDraw(strategyId, ids)
	return
	//扣减库存

}

func (db *DrawBase) queryExcludeAwardIds(id int64) []gormModel.StrategyDetail {
	//返回没有库存的奖品id
	//gorm 返回strategyDetail 中对应的wardIds
	return db.rep.QueryExcludeAwardIds(id)
}

func (db *DrawBase) deductStock(strategyId int32, awardId string) (ret string) {
	id, _ := strconv.Atoi(awardId)
	awardid := int64(id)
	//return
	if db.rep.DeductStock(int64(strategyId), awardid) != nil {
		ret = awardId
		return
	}
	return
}

func (db *DrawBase) buildDrawResult(uid string, strategyId int64, awardId string, strategy vo.StrategyBriefVO) *res.DrawResult {
	if "" == awardId {
		//未抽奖
		db.log.Infof("执行策略抽奖完成【未中奖】 用户: %v, 策略id: %v ", uid, strategyId)
		return &res.DrawResult{
			Uid:        uid,
			StrategyId: strategyId,
			DrawResult: data.FAIL,
		}
	}
	awardid, _ := strconv.Atoi(awardId)
	award, _ := db.rep.QueryAwardInfoByAwardId(int64(awardid))
	awardStr := strconv.Itoa(int(award.AwardId))
	drawAward := vo.NewDrawAwardVo(uid, awardStr, award.AwardType, award.AwardName, award.AwardContent, strategy.StrategyMode, strategy.GrantType, strategy.GrandDate)
	return &res.DrawResult{
		Uid:           uid,
		StrategyId:    strategyId,
		DrawResult:    data.SUCESS,
		DrawAwardInfo: *drawAward,
	}

}
