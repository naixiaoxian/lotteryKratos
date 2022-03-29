package gormModel

import "time"

type Award struct {
	Id           int64     `gorm:"Column:id;type:bigint"`
	AwardId      int64     `gorm:"Column:awardId;type:bigint"`
	AwardType    int       `gorm:"Column:awardType;type:int"`
	AwardCount   int       `gorm:"Column:awardCount;type:int"`
	AwardName    string    `gorm:"Column:awardName;type:varchar(64)"`
	AwardContent string    `gorm:"Column:awardContent;type:varchar(128)"`
	CreateTime   time.Time `gorm:"Column:createTime;type:datetime"`
	UpdateTime   time.Time `gorm:"Column:updateTime;type:datetime"`
}

func (*Award) TableName() string {
	return "award"
}
