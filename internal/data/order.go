package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"lotteryKratos/internal/biz/award/goods"
	"lotteryKratos/internal/data/gormModel"
	"time"
)

func NewOrderRep(data *Data, logger log.Logger) goods.AwardRepositoryImpl {
	return &OrderRep{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type OrderRep struct {
	data *Data
	log  *log.Helper
}

func (or *OrderRep) UpdateUserAwardState(uid string, orderId int64, awardId string, grantState int) {
	model := &gormModel.UserStrategyExport{
		Uid:          uid,
		ActivityId:   0,
		OrderId:      orderId,
		StrategyId:   0,
		StrategyType: 0,
		GrantType:    0,
		GrantDate:    time.Now(),
		GrantState:   grantState,
		AwardId:      awardId,
		AwardType:    0,
		AwardName:    "",
		AwardContent: "",
		UUid:         "",
		Mqstate:      0,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	or.data.Db.Create(model)
	return
}
