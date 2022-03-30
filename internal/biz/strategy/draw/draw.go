package draw

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	algorithm2 "lotteryKratos/internal/biz/strategy/algorithm"
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

type StrategyRepoImpl interface {
	QueryExcludeAwardIds(id int64) []gormModel.StrategyDetail
	DeductStock(strategyId int64, awardId int64) error
	QueryAwardInfoByAwardId(id int64) (gormModel.Award, error)
	//获取rich聚合根
	QueryStrategyRich(id int64) (aggregates.StrategyRich, error)
}
type DrawBase struct {
	rep StrategyRepoImpl
	log *log.Helper
	singleDraw algorithm2.SingleRateRandomDrawAlgorithm
	entireDraw algorithm2.EntiretyRateRandomDrawAlgorithm
}

func NewDraBaseDomain(repo StrategyRepoImpl, logger log.Logger,
	singleDraw algorithm2.SingleRateRandomDrawAlgorithm,
	entireDraw algorithm2.EntiretyRateRandomDrawAlgorithm) *DrawBase {
	return &DrawBase{rep: repo, log: log.NewHelper(logger), singleDraw: singleDraw, entireDraw: entireDraw}
}

func (db *DrawBase) getAlgorithm(mode int) algorithm2.DrawImpl {
	if mode == 1 {
		return &db.entireDraw
	}
	return &db.singleDraw
}

func (db *DrawBase) DoDrawExec(req req.DrawReq) (ret *res.DrawResult) {
	//获取策略模式
	stragetyRich, _ := db.rep.QueryStrategyRich(req.StrategyId)
	strategyBrief := stragetyRich.StrategyBriefVo
	db.checkAndInitRateData(req.StrategyId, strategyBrief.StrategyMode, stragetyRich.StrategyDetails)
	//抽奖算法
	exludeAwardIds := db.rep.QueryExcludeAwardIds(req.StrategyId)
	fmt.Println("ExecAwardIDs", exludeAwardIds)
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

func (db *DrawBase) drawAlgorithm(strategyId int64, drawImpl algorithm2.DrawImpl, strategyDetailList []gormModel.StrategyDetail) (awardId string) {
	//strategyRich :=
	ids := make([]string, 0)
	for _, strategyDetail := range strategyDetailList {
		ids = append(ids, strconv.Itoa(strategyDetail.AwardId))
	}
	awardId = drawImpl.RandomDraw(strategyId, ids)
	//return
	if "" == awardId || "未中奖" == awardId {
		awardId = ""
		return
	}
	isSuccess := db.deductStock(strategyId, awardId)
	return isSuccess
	//扣减库存

}

func (db *DrawBase) deductStock(strategyId int64, awardId string) (ret string) {
	id, _ := strconv.Atoi(awardId)
	awardid := int64(id)
	//return
	rets := db.rep.DeductStock(strategyId, awardid)
	if rets == nil {
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
			DrawResult: 0, //data.fail
		}
	}
	awardid, _ := strconv.Atoi(awardId)
	award, _ := db.rep.QueryAwardInfoByAwardId(int64(awardid))
	awardStr := strconv.Itoa(int(award.AwardId))
	drawAward := vo.NewDrawAwardVo(uid, awardStr, award.AwardType, award.AwardName, award.AwardContent, strategy.StrategyMode, strategy.GrantType, strategy.GrandDate)
	db.log.Infof("执行策略抽奖完成【中奖】 用户: %v, 策略id: %v 奖品：%v ,id %v", uid, strategyId, drawAward, awardid)
	return &res.DrawResult{
		Uid:           uid,
		StrategyId:    strategyId,
		DrawResult:    1, //data.success
		DrawAwardInfo: *drawAward,
	}

}

