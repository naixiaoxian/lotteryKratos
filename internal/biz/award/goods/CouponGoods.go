package goods

import (
	"lotteryKratos/internal/data/award/req"
	"lotteryKratos/internal/data/award/res"
)

type CouponGoods struct {
	DB *DistributionBase
}

func (cg *CouponGoods) DoDistribution(req req.GoodsReq) (db *res.DistributionRes) {
	//1 优惠券
	cg.DB.updateUserAwardState(req.Uid, req.OrderId, req.AwardId, 1)
	return &res.DistributionRes{
		Uid:  req.Uid,
		Code: 1,
		Info: "优惠券发送成功",
	}
}
