package vo

import "time"

// ActivityBillVO 活动账单【库存、状态、日期、个人参与次数】
type ActivityVO struct {
	Id                int64
	ActivityId        int64
	ActivityName      string
	ActivityDesc      string
	BeginDateTime     time.Time
	EndDateTime       time.Time
	StockCount        int
	StockSurplusCount int
	State             int // 活动状态 1 编辑2提交审核3 撤销审核4审核通过5 开始运行 6 拒绝 7 关闭 8 开启
	StrategyId        int64
	TakeCount         int
	UserTakeLeftCount int
	Creator           string
	CreateTime        time.Time
	UpdateTime        time.Time
}
