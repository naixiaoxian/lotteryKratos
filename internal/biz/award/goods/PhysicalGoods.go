package goods

import (
	"lotteryKratos/internal/data/award/req"
	"lotteryKratos/internal/data/award/res"
)

type PhysicalGoods struct {
	DB *DistributionBase
}

func (cg *PhysicalGoods) DoDistribution(req req.GoodsReq) (db *res.DistributionRes) {
	//3 物理奖品发放成功
	cg.DB.updateUserAwardState(req.Uid, req.OrderId, req.AwardId, 1)
	return &res.DistributionRes{
		Uid:  req.Uid,
		Code: 1,
		Info: "实物商品发送成功",
	}
}
