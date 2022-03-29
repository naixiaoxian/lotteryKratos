package gormModel

import "time"

type Award struct {
	Id           int64     `gorm:"Column:id;type:bigint"`
	AwardId      int64     `gorm:"Column:award_id;type:bigint"`
	AwardType    int       `gorm:"Column:award_type;type:int"`
	AwardCount   int       `gorm:"Column:award_count;type:int"`
	AwardName    string    `gorm:"Column:award_name;type:varchar(64)"`
	AwardContent string    `gorm:"Column:award_content;type:varchar(128)"`
	CreateTime   time.Time `gorm:"Column:create_time;type:datetime"`
	UpdateTime   time.Time `gorm:"Column:update_time;type:datetime"`
}

func (*Award) TableName() string {
	return "award"
}
