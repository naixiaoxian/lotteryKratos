package goods

type DistributionBase struct {
	Impl AwardRepositoryImpl
}

type AwardRepositoryImpl interface {
	UpdateUserAwardState(uid string, orderId int64, awardId string, grantState int)
}

func (db *DistributionBase) updateUserAwardState(uid string, orderId int64, awardId string, grantState int) {
	db.Impl.UpdateUserAwardState(uid, orderId, awardId, grantState)
}
