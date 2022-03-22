package vo

type AwardRateVo struct {
	AwardId   string
	AwardRate float32
}

func NewAwardRateVo(awardId string, awardRate float32) (arv *AwardRateVo) {
	arv.AwardId = awardId
	arv.AwardRate = awardRate
	return
}

func (arv *AwardRateVo) GetAwardRate() float32 {
	return arv.AwardRate
}
