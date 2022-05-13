package stateFlow

import (
	"lotteryKratos/internal/biz/activity"
	"lotteryKratos/internal/common"
)

type PassState struct {
	handler *StateHandler
}

func NewPassState(rep activity.IActivityRepImpl) (ret StateHandlerImp) {
	return &PassState{NewStateHandler(rep)}
}

func (a PassState) Arraignment(i int, i2 int) (ret string) {
	return "已审核状态不可重复提审"
}

func (a PassState) CheckPass(i int, i2 int) (ret string) {
	return "已审核状态不可重复审核"
}

func (a PassState) CheckRefuse(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_REFUSE)
	if success {
		ret = "活动审核拒绝完成"
	} else {
		ret = "活动状态变更失败"
	}
	return
}

func (a PassState) CheckRevoke(i int, i2 int) (ret string) {
	return "审核通过不可撤销"
}

func (a PassState) Close(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_CLOSE)
	if success {
		ret = "活动审核关闭完成"
	} else {
		ret = "活动状态变更失败"
	}
	return
}

func (a PassState) Open(i int, i2 int) (ret string) {
	return "非关闭活动不可开启"
}

func (a PassState) Doing(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_DOING)
	if success {
		ret = "活动审核启动完成"
	} else {
		ret = "活动状态变更失败"
	}
	return
}
