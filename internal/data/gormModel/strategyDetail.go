package gormModel

import (
	"github.com/shopspring/decimal"
	"time"
)

type StrategyDetail struct {
	Id                int64           `gorm:"Column:id;type:bigint"`
	StrategyId        int64           `gorm:"Column:strategy_id;type:int"`
	AwardId           int             `gorm:"Column:award_id;type:int"`
	AwardCount        int             `gorm:"Column:award_count;type:int"`
	AwardRate         decimal.Decimal `gorm:"Column:award_rate;type:decimal(5,2)"`
	AwardSurplusCount int64           `gorm:"Column:award_surplusCount;type:bigint"`
	AwardDesc         string          `gorm:"Column:award_desc;type:varchar(255)"`
	CreateTime        time.Time       `gorm:"Column:Create_time;type:datetime"`
	UpdateTime        time.Time       `gorm:"Column:Update_time;type:datetime"`
}

func (*StrategyDetail) TableName() string {
	return "strategy_detail"
}
