package partake

import (
	"fmt"
	"lotteryKratos/internal/biz/activity"
	"lotteryKratos/internal/common"
	"lotteryKratos/internal/data/activity/req"
	"lotteryKratos/internal/data/activity/res"
	"lotteryKratos/internal/data/activity/vo"
	"strconv"
	"time"
)

type Partake interface {
	DoPartake(req req.PartakeReq) (ret res.PartakeResult, errInfo string)
	RecordDrawOrder(drawOrder vo.DrawOrderVO) string
	UpdateInvoiceMqState(uId string, orderId int64, mqState int)
	ScanInvoiceMqState() []vo.InvoiceVO //后期做成分库分表。前期使用单库单表做查询
	UpdateActivityStock(vo vo.ActivityPartakeRecordVO)
}

type IUserTakeActivityRepository interface {
	BeginTrans() //用于开启事务
	Commit()
	Rollback() //事务相关在这里实现
	SubtractionLeftCount(activityId int64, activityName string, takeCount int, userTakeLeftcount int, uId string) int
	TakeActivity(activityId int64, activityName string, strategyId int64, takeCount int, userTakeLeftCount int, uId string, takeDate time.Time, takeId int64)
	LockTackActivity(uId string, activityId int64, takeId int64)
	SaveUserStrategyExport(orderVO vo.DrawOrderVO)
	QueryNoConsumedTakeActivityOrder(activityId int64, uId string) vo.UserTakeActivityVO
	UpdateInvoiceMqState(uId string, orderId int64, mqState int)
	ScanInvoiceMqState() []vo.InvoiceVO
	UpdateActivityStock(recordVO vo.ActivityPartakeRecordVO)
}

type BasePartake struct {
	iUserTakeImpl IUserTakeActivityRepository
	iActivityImpl activity.IActivityRepImpl
}

func (BP *BasePartake) DoPartake(req req.PartakeReq) (ret res.PartakeResult, errInfo string) {
	//查询是否存在未执行抽奖的活动单
	userTakeActivityVo := BP.iUserTakeImpl.QueryNoConsumedTakeActivityOrder(req.ActivityId, req.Uid)
	if userTakeActivityVo != (vo.UserTakeActivityVO{}) {
		ret = res.PartakeResult{
			StrategyId:        userTakeActivityVo.StrategyId,
			TakeId:            userTakeActivityVo.TakeId,
			StockCount:        0,
			StockSurplusCount: 0,
		}
		return
	}
	//查询活动账单
	activityBillVO := BP.iActivityImpl.QueryActivityBill(req)
	//活动账单校对
	errInfo = BP.CheckActivityBill(req, activityBillVO)
	if "" != errInfo {
		return
	}
	//扣减活动库存
	subActivityResult := BP.SubtractionActivityStockByRedis(req.Uid, req.ActivityId, activityBillVO.StockCount)
	if 1 != 1 {
		//如果报错那么加回去
		errInfo = "活动库存扣减失败"
		BP.RecoverActivityCacheStockByRedis(req.ActivityId, subActivityResult.StockKey, "活动库存扣减失败")
		return
	}
	//插入活动信息
	takeId, _ := strconv.Atoi(time.Now().String())
	grabResult := BP.GrabActivity(req, activityBillVO, int64(takeId))
	if "" != grabResult {
		errInfo = grabResult
		BP.RecoverActivityCacheStockByRedis(req.ActivityId, subActivityResult.StockKey, errInfo)
		return
	}
	//扣减活动库存
	BP.RecoverActivityCacheStockByRedis(req.ActivityId, subActivityResult.StockKey, "")
	ret = res.PartakeResult{
		StrategyId:        activityBillVO.StrategyId,
		TakeId:            int64(takeId),
		StockCount:        activityBillVO.StockCount,
		StockSurplusCount: subActivityResult.StockSurplusCount,
	}
	return
	//返回结果
}

func (BP *BasePartake) CheckActivityBill(partakeReq req.PartakeReq, billVO vo.ActivityBillVO) string {
	if common.ACTIVITY_STATE_DOING != billVO.State {
		return "当前活动状态非可用"
	}
	if billVO.BeginDateTime.After(partakeReq.PartakeDate) || billVO.EndDateTime.Before(partakeReq.PartakeDate) {
		return "活动时间范围非可用"
	}
	if billVO.StockSurplusCount <= 0 {
		return "活动剩余库存非可用"
	}
	if billVO.UserTakeLeftCount <= 0 {
		return "个人领取次数非可用"
	}
	return ""
}

func (BP *BasePartake) SubtractionActivityStockByRedis(uId string, activityId int64, stockCount int) res.StockResult {
	return BP.iActivityImpl.SubtractionActivityStockByRedis(uId, activityId, stockCount)
}

func (BP *BasePartake) RecoverActivityCacheStockByRedis(activityId int64, tokenKey string, code string) {
	BP.iActivityImpl.RecoverActivityCacheStockByRedis(activityId, tokenKey, code)
}

func (BP *BasePartake) GrabActivity(partakeReq req.PartakeReq, billVO vo.ActivityBillVO, takeId int64) string {
	updateCount := BP.iUserTakeImpl.SubtractionLeftCount(billVO.ActivityId, billVO.ActivityName, billVO.TakeCount, billVO.UserTakeLeftCount, partakeReq.Uid)
	if updateCount == 0 {
		fmt.Println("领取活动失败", partakeReq, billVO, takeId)
		return "更新数量失败，领取失败"
	}
	BP.iUserTakeImpl.TakeActivity(billVO.ActivityId, billVO.ActivityName, billVO.StrategyId, billVO.TakeCount, billVO.UserTakeLeftCount, partakeReq.Uid, partakeReq.PartakeDate, takeId)
	return ""
}
