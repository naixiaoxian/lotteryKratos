package gormModel

import "time"

type Strategy struct {
	Id           int64
	StrategyId   int64
	StrategyDesc string
	StrategyMode int
	GrantType    int
	GrantDate    time.Time
	ExtraInfo    string
	CreateTime   time.Time
	UpdateTime   time.Time
}

func (*Strategy) TableName() string {
	return "strategy"
}
