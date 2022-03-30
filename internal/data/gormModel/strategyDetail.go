package gormModel

import (
	"time"
)

type StrategyDetail struct {
	Id                int64   `gorm:"Column:id;type:bigint"`
	StrategyId        int64   `gorm:"Column:strategy_id;type:int"`
	AwardId           int     `gorm:"Column:award_id;type:int"`
	AwardName         string  `gorm:"Column:award_name;type:varchar(255)"`
	AwardCount        int     `gorm:"Column:award_count;type:int"`
	AwardRate         float32 `gorm:"Column:award_rate;type:decimal(5,2)"`
	AwardSurplusCount int     `gorm:"Column:award_surplus_count;type:bigint"`
	//AwardDesc         string    `gorm:"Column:award_desc;type:varchar(255)"`
	CreateTime time.Time `gorm:"Column:Create_time;type:datetime"`
	UpdateTime time.Time `gorm:"Column:Update_time;type:datetime"`
}

func (*StrategyDetail) TableName() string {
	return "strategy_detail"
}
