package award

import "lotteryKratos/internal/biz/award/goods"

type GoodsConfig struct {
	couponGoods   *goods.CouponGoods
	descGoods     *goods.DescGoods
	physicalGoods *goods.PhysicalGoods
	redeemCode    *goods.RedeemCodeGoods
	Goods         map[int]goods.IDistributionGoods
}

func NewGoodsConfig(
	couponGoods *goods.CouponGoods,
	descGoods *goods.DescGoods,
	physicalGoods *goods.PhysicalGoods,
	redeemCode *goods.RedeemCodeGoods,
) (config GoodsConfig) {
	config.couponGoods = couponGoods
	config.descGoods = descGoods
	config.physicalGoods = physicalGoods
	config.redeemCode = redeemCode
	goodMap := make(map[int]goods.IDistributionGoods)
	//1 优惠券 2 描述 3 物理发货 4 优惠券
	goodMap[1] = config.couponGoods
	goodMap[2] = config.descGoods
	goodMap[3] = config.physicalGoods
	goodMap[4] = config.redeemCode
	config.Goods = goodMap
	return
}

func (gc *GoodsConfig) GetDistributionGoodsService(awardType int) (ret goods.IDistributionGoods) {
	ret, _ = gc.Goods[awardType]
	return ret
}
