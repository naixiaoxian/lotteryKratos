package gormModel

import (
	"google.golang.org/genproto/googleapis/type/decimal"
	"time"
)

type StrategyDetail struct {
	Id         int64
	StrategyId int64
	AwardId    int64
	AwardCount int
	AwardRate  decimal.Decimal `gorm:"type:decimal(5,2)"`
	CreateTime time.Time
	UpdateTime time.Time
}

func (*StrategyDetail) TableName() string {
	return "strategy_detail"
}
