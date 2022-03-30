package gormModel

import "time"

type Strategy struct {
	Id           int64
	StrategyId   int64     `gorm:"Column:strategy_id;type:bigint"`
	StrategyDesc string    `gorm:"Column:strategy_desc;type:varchar(128)"`
	StrategyMode int       `gorm:"Column:strategy_mode;type:int"`
	GrantType    int       `gorm:"Column:grant_type;type:int"`
	GrantDate    time.Time `gorm:"Column:grant_date;type:datetime"`
	ExtInfo      string    `gorm:"Column:ext_info;type:varchar(128)"`
	CreateTime   time.Time `gorm:"Column:create_time;type:datetime"`
	UpdateTime   time.Time `gorm:"Column:update_time;type:datetime"`
}

func (*Strategy) TableName() string {
	return "strategy"
}
