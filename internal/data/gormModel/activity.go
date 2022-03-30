package gormModel

import "time"

type Activity struct {
	Id                int64
	ActivityID        int
	ActivityName      string
	ActivityDesc      string
	BeginDateTime     time.Time
	EndDateTime       time.Time
	StockCount        int
	TakeCount         int
	StockSurplusCount int
	State             int
	Creator           string
	CreateTime        time.Time
	UpdateTime        time.Time
}

func (*Activity) TableName() string {
	return "activity"
}
