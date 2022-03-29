package goods

import (
	"lotteryKratos/internal/data/award/req"
	"lotteryKratos/internal/data/award/res"
)

type DescGoods struct {
	DB *DistributionBase
}

func (cg *DescGoods) DoDistribution(req req.GoodsReq) (db *res.DistributionRes) {
	//2 优惠券
	cg.DB.updateUserAwardState(req.Uid, req.OrderId, req.AwardId, 1)
	return &res.DistributionRes{
		Uid:  req.Uid,
		Code: 1,
		Info: "描述商品发送成功",
	}
}
