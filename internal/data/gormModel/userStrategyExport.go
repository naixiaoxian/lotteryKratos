package gormModel

import "time"

type UserStrategyExport struct {
	Id           int64     `gorm:"Column:id;type:bigint"`
	Uid          string    `gorm:"Column:uId;type:longtext"`
	ActivityId   int64     `gorm:"Column:activityId;type:bigint"`
	OrderId      int64     `gorm:"Column:orderId;type:bigint"`
	StrategyId   int64     `gorm:"Column:strategyId;type:bigint"`
	StrategyType int       `gorm:"Column:strategyType;type:int"`
	GrantType    int       `gorm:"Column:grantType;type:int"`
	GrantDate    time.Time `gorm:"Column:GrantDate;type:timestamp"`
	GrantState   int       `gorm:"Column:GrantState;type:int"`
	AwardId      string    `gorm:"Column:AwardId;type:bigint"`
	AwardType    int       `gorm:"Column:AwardType;type:int"`
	AwardName    string    `gorm:"Column:AwardName;type:longtext"`
	AwardContent string    `gorm:"Column:AwardContent;type:longtext"`
	UUid         string    `gorm:"Column:UUid;type:longtext"`
	Mqstate      int       `gorm:"Column:mqstate;type:int"`
	CreateTime   time.Time `gorm:"Column:createTime;type:timestamp"`
	UpdateTime   time.Time `gorm:"Column:updateTime;type:timestamp"`
}

func (*UserStrategyExport) TableName() string {
	return "user_strategy_export"
}
