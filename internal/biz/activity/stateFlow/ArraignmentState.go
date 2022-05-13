package stateFlow

import (
	"lotteryKratos/internal/biz/activity"
	"lotteryKratos/internal/common"
)

type ArraignmentState struct {
	handler *StateHandler
}

func NewArraignmentState(rep activity.IActivityRepImpl) (ret StateHandlerImp) {
	return &ArraignmentState{NewStateHandler(rep)}
}

func (a ArraignmentState) Arraignment(i int, i2 int) (ret string) {
	return "待审核状态不可重复提审"
}

func (a ArraignmentState) CheckPass(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_PASS)
	if success {
		ret = "活动审核完成"
	} else {
		ret = "活动状态变更失败"
	}
	return
}

func (a ArraignmentState) CheckRefuse(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_REFUSE)
	if success {
		ret = "活动审核拒绝完成"
	} else {
		ret = "活动状态变更失败"
	}
	return
}

func (a ArraignmentState) CheckRevoke(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_EDIT)
	if success {
		ret = "活动审核撤销回到编辑中"
	} else {
		ret = "活动状态变更失败"
	}
	return
}

func (a ArraignmentState) Close(i int, i2 int) (ret string) {
	ret = "非拒绝活动不可关闭"
	return
}

func (a ArraignmentState) Open(i int, i2 int) (ret string) {
	ret = "非关闭活动不可开启"
	return
}

func (a ArraignmentState) Doing(i int, i2 int) (ret string) {
	ret = "待审核活动不可执行活动中变更"
	return
}
