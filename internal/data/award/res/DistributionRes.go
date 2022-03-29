package res

type DistributionRes struct {
	Uid         string
	Code        int
	Info        string
	statementId string
}

func (d *DistributionRes) newRes(uid string, code int, info string) *DistributionRes {
	return &DistributionRes{
		Uid:         uid,
		Code:        code,
		Info:        info,
		statementId: "",
	}
}
