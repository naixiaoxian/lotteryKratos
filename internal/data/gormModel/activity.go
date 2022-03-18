package gormModel

import "time"

type ActivityModel struct {
	Id            int64
	ActivityName  string
	ActivityDesc  string
	BeginDateTime time.Time
	EndDateTime   time.Time
	StockCount    int32
	TakeCount     int32
	State         int32
	Creator       string
	CreateTime    time.Time
	UpdateTime    time.Time
}

func (*ActivityModel) TableName() string {
	return "activity"
}
