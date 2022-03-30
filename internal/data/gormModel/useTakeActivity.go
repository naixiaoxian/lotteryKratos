package gormModel

import "time"

type UserTakeActivity struct {
	Id           int64     `gorm:"Column:id;type:bigint"`
	UId          string    `gorm:"Column:uId;type:text"`
	TakeId       int64     `gorm:"Column:takeId;type:bigint"`
	ActivityId   int64     `gorm:"Column:activityId;type:bigint"`
	ActivityName int       `gorm:"Column:activityName;type:text"`
	TakeDate     int       `gorm:"Column:takeDate;type:timestamp"`
	TakeCount    time.Time `gorm:"Column:takeCount;type:int"`
	Uuid         int       `gorm:"Column:uuid;type:text"`
	CreateTime   time.Time `gorm:"Column:createTime;type:timestamp"`
	UpdateTime   time.Time `gorm:"Column:updateTime;type:timestamp"`
}

func (*UserTakeActivity) TableName() string {
	return "user_take_activity"
}
