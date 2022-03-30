package vo

import "lotteryKratos/internal/data/award/vo"

type InvoiceVO struct {
	Uid             string
	OrderId         int64
	AwardId         string
	AwardType       int
	AwardName       string
	AwardContent    string
	ShippingAddress vo.ShippingAddress
	ExtInfo         string
}
