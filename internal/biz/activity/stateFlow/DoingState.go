package stateFlow

import (
	"lotteryKratos/internal/biz/activity"
	"lotteryKratos/internal/common"
)

type DoingState struct {
	handler *StateHandler
}

func NewDoingState(rep activity.IActivityRepImpl) (ret StateHandlerImp) {
	return &DoingState{NewStateHandler(rep)}
}

func (a DoingState) Arraignment(i int, i2 int) (ret string) {
	return "活动中不可提审"
}

func (a DoingState) CheckPass(i int, i2 int) (ret string) {
	return "活动中不可审核通过"
}

func (a DoingState) CheckRefuse(i int, i2 int) (ret string) {
	return "活动中不可审核拒绝"
}

func (a DoingState) CheckRevoke(i int, i2 int) (ret string) {
	return "活动中不可撤销审核"
}

func (a DoingState) Close(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_CLOSE)
	if success {
		ret = "活动关闭完成"
	} else {
		ret = "活动状态变更失败"
	}
	return
}

func (a DoingState) Open(i int, i2 int) (ret string) {
	return "非关闭活动不可开启"
}

func (a DoingState) Doing(i int, i2 int) (ret string) {
	return "活动中不可重复执行"
}
