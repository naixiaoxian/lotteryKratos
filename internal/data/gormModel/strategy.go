package gormModel

import "time"

type Strategy struct {
	Id           int64
	StrategyId   int64     `gorm:"Column:strategyId;type:bigint"`
	StrategyDesc string    `gorm:"Column:strategyDesc;type:varchar(128)"`
	StrategyMode int       `gorm:"Column:strategyMode;type:int"`
	GrantType    int       `gorm:"Column:grantType;type:int"`
	GrantDate    time.Time `gorm:"Column:grantDate;type:datetime"`
	ExtraInfo    string    `gorm:"Column:extraInfo;type:varchar(128)"`
	CreateTime   time.Time `gorm:"Column:createTime;type:datetime"`
	UpdateTime   time.Time `gorm:"Column:updateTime;type:datetime"`
}

func (*Strategy) TableName() string {
	return "strategy"
}
