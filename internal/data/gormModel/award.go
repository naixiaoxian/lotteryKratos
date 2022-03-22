package gormModel

import "time"

type Award struct {
	Id           int64
	AwardId      int64
	AwardType    int
	AwardCount   int
	AwardName    string
	AwardContent string
	CreateTime   time.Time
	UpdateTime   time.Time
}

func (*Award) TableName() string {
	return "award"
}
