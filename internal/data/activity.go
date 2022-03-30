package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"lotteryKratos/internal/biz"
	"lotteryKratos/internal/data/gormModel"
	"time"
)

type activityRepo struct {
	data *Data
	log  *log.Helper
}

func NewActivityRepo(data *Data, logger log.Logger) biz.ActivityRepoImpl {
	return &activityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *activityRepo) List(ctx context.Context) (rv []*biz.Activity, err error) {
	//todo 注意page 跟pagesize 的分页
	fmt.Println(ctx)
	//对应的query 查询方法与对应方式需要新增一个activityQuery类 去解决相关查询的问题
	//这个查询需要重新写在activity query 方法中
	activityDB := ar.data.Db.Model(&gormModel.Activity{})
	var count int64
	activityDB.Count(&count)
	page := 1
	pageSize := 5
	var activityList []gormModel.Activity
	err = activityDB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&activityList).Error
	if err != nil {
		return
	}
	for _, p := range activityList {
		ac := &biz.Activity{}
		rv = append(rv, ac.ConvertFrom(p))
	}
	return
}

func (ar *activityRepo) Get(ctx context.Context, id int64) (rv *biz.Activity, err error) {
	fmt.Println(ctx)
	fmt.Println("get", id)
	var activity gormModel.Activity
	//err = ar.data.Db.First(&activity, id).Error
	err = ar.data.Db.Where("id = ?", id).First(&activity).Error
	fmt.Println(activity)
	//rc :=
	//rv.ConvertFrom(activity)
	rv = &biz.Activity{
		ID:            activity.Id,
		Name:          activity.ActivityName,
		Desc:          activity.ActivityDesc,
		BeginDateTime: activity.BeginDateTime,
		EndDateTime:   activity.EndDateTime,
		StockCount:    int32(activity.StockCount),
		TakeCount:     int32(activity.TakeCount),
		State:         int32(activity.State),
		Creator:       activity.Creator,
		CreateTime:    activity.CreateTime,
		UpdateTime:    activity.UpdateTime,
	}
	if err != nil {
		return
	}
	return
}

func (ar *activityRepo) Create(ctx context.Context, article *biz.Activity) (rv *biz.Activity, err error) {
	fmt.Println(ctx)
	model := gormModel.Activity{
		ActivityName:  article.Name,
		ActivityDesc:  article.Desc,
		BeginDateTime: article.BeginDateTime,
		EndDateTime:   article.EndDateTime,
		StockCount:    int(article.StockCount),
		TakeCount:     0,
		State:         0,
		Creator:       article.Creator,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	err = ar.data.Db.Create(&model).Error
	rv = article.ConvertFrom(model)
	return
}

func (ar *activityRepo) Update(ctx context.Context, id int64, article *biz.Activity) (rv *biz.Activity, err error) {
	fmt.Println(ctx, id, article)
	return
}

func (ar *activityRepo) Delete(ctx context.Context, id int64) (err error) {
	//todo 注意page 跟pagesize 的分页
	fmt.Println(ctx)
	err = ar.data.Db.Where("id = ?", id).Delete(&gormModel.Activity{}).Error
	return
}
