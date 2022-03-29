package gormModel

import (
	"github.com/shopspring/decimal"
	"time"
)

type StrategyDetail struct {
	Id                int64           `gorm:"Column:id;type:bigint"`
	StrategyId        int64           `gorm:"Column:strategyId;type:int"`
	AwardId           int             `gorm:"Column:awardId;type:int"`
	AwardCount        int             `gorm:"Column:awardCount;type:int"`
	AwardRate         decimal.Decimal `gorm:"Column:awardRate;type:decimal(5,2)"`
	AwardSurplusCount int64           `gorm:"Column:awardSurplusCount;type:bigint"`
	AwardDesc         string          `gorm:"Column:awardDesc;type:varchar(255)"`
	CreateTime        time.Time       `gorm:"Column:CreateTime;type:datetime"`
	UpdateTime        time.Time       `gorm:"Column:UpdateTime;type:datetime"`
}

func (*StrategyDetail) TableName() string {
	return "strategy_detail"
}
