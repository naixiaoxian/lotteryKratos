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
	activityDB := ar.data.db.Model(&gormModel.ActivityModel{})
	var count int64
	activityDB.Count(&count)
	page := 1
	pageSize := 5
	var activityList []gormModel.ActivityModel
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
	var activity gormModel.ActivityModel
	//err = ar.data.db.First(&activity, id).Error
	err = ar.data.db.Where("id = ?", id).First(&activity).Error
	fmt.Println(activity)
	//rc :=
	//rv.ConvertFrom(activity)
	rv = &biz.Activity{
		ID:            activity.Id,
		Name:          activity.ActivityName,
		Desc:          activity.ActivityDesc,
		BeginDateTime: activity.BeginDateTime,
		EndDateTime:   activity.EndDateTime,
		StockCount:    activity.StockCount,
		TakeCount:     activity.TakeCount,
		State:         activity.State,
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
	model := gormModel.ActivityModel{
		ActivityName:  article.Name,
		ActivityDesc:  article.Desc,
		BeginDateTime: article.BeginDateTime,
		EndDateTime:   article.EndDateTime,
		StockCount:    article.StockCount,
		TakeCount:     0,
		State:         0,
		Creator:       article.Creator,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	err = ar.data.db.Create(&model).Error
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
	err = ar.data.db.Where("id = ?", id).Delete(&gormModel.ActivityModel{}).Error
	return
}

//type ActivityRepoImpl interface {
//	List(ctx context.Context) ([]*Activity, error)
//	Get(ctx context.Context, id int64) (*Activity, error)
//	Create(ctx context.Context, article *Activity) (*Activity, error)
//	Update(ctx context.Context, id int64, article *Activity) (*Activity, error)
//	Delete(ctx context.Context, id int64) error
//}
