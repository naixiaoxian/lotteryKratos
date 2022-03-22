package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"lotteryKratos/internal/data/gormModel"
	"time"
)

type Activity struct {
	ID            int64
	Name          string
	Desc          string
	BeginDateTime time.Time
	EndDateTime   time.Time
	StockCount    int32
	TakeCount     int32
	State         int32
	Creator       string
	CreateTime    time.Time
	UpdateTime    time.Time
}

//TODO 这个方法有点问题。需要重新返回一个逻辑
func (*Activity) ConvertFrom(model gormModel.Activity) (ac *Activity) {
	ac.ID = model.Id
	ac.Name = model.ActivityName
	ac.Desc = model.ActivityDesc
	ac.BeginDateTime = model.BeginDateTime
	ac.EndDateTime = model.EndDateTime
	ac.StockCount = model.StockCount
	ac.TakeCount = model.TakeCount
	ac.State = model.State
	ac.Creator = model.Creator
	ac.CreateTime = model.CreateTime
	ac.UpdateTime = model.UpdateTime
	return ac
}

type ActivityRepoImpl interface {
	List(ctx context.Context) ([]*Activity, error)
	Get(ctx context.Context, id int64) (*Activity, error)
	Create(ctx context.Context, article *Activity) (*Activity, error)
	Update(ctx context.Context, id int64, article *Activity) (*Activity, error)
	Delete(ctx context.Context, id int64) error
}

type ActivityDomain struct {
	impl ActivityRepoImpl
	log  *log.Helper
}

func NewActivityDomain(repo ActivityRepoImpl, logger log.Logger) *ActivityDomain {
	return &ActivityDomain{impl: repo, log: log.NewHelper(logger)}
}

// List domain 调用抽象来实现具体repo 行为
func (ar *ActivityDomain) List(ctx context.Context) (ret []*Activity, err error) {
	ret, err = ar.impl.List(ctx)
	if err != nil {
		return
	}
	return
}
func (ar *ActivityDomain) Get(ctx context.Context, id int64) (ret *Activity, err error) {
	ret, err = ar.impl.Get(ctx, id)
	if err != nil {
		return
	}
	return
}
func (ar *ActivityDomain) Create(ctx context.Context, activity *Activity) (ret *Activity, err error) {
	ret, err = ar.impl.Create(ctx, activity)
	if err != nil {
		return
	}
	return
}
func (ar *ActivityDomain) Update(ctx context.Context, id int64, activity *Activity) (ret *Activity, err error) {
	ret, err = ar.impl.Update(ctx, id, activity)
	if err != nil {
		return
	}
	return
}
func (ar *ActivityDomain) Delete(ctx context.Context, id int64) (err error) {
	err = ar.impl.Delete(ctx, id)
	if err != nil {
		return
	}
	return
}
