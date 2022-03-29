package goods

import (
	"lotteryKratos/internal/data/award/req"
	"lotteryKratos/internal/data/award/res"
)

type IDistributionGoods interface {
	DoDistribution(req req.GoodsReq) *res.DistributionRes
}
