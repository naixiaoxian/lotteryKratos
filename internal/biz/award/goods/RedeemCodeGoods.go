package goods

import (
	"lotteryKratos/internal/data/award/req"
	"lotteryKratos/internal/data/award/res"
)

type RedeemCodeGoods struct {
	DB *DistributionBase
}

func (cg *RedeemCodeGoods) DoDistribution(req req.GoodsReq) (db *res.DistributionRes) {
	//4 兑换券发送成功
	cg.DB.updateUserAwardState(req.Uid, req.OrderId, req.AwardId, 1)
	return &res.DistributionRes{
		Uid:  req.Uid,
		Code: 1,
		Info: "兑换券商品发送成功",
	}
}
