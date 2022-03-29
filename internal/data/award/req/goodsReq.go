package req

import "lotteryKratos/internal/data/award/vo"

type GoodsReq struct {
	Uid             string
	OrderId         int64
	AwardId         string
	AwardName       string
	AwardContent    string
	ShippingAddress vo.ShippingAddress
	ExtInfo         string
}

func (*GoodsReq) newGoodsReq(uid string, orderId int64, awardId string, awdName string, content string) *GoodsReq {
	return &GoodsReq{
		Uid:             uid,
		OrderId:         orderId,
		AwardId:         awardId,
		AwardName:       awdName,
		AwardContent:    content,
		ShippingAddress: vo.ShippingAddress{},
		ExtInfo:         "",
	}
}
