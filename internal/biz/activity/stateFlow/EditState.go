package stateFlow

import (
	"lotteryKratos/internal/biz/activity"
	"lotteryKratos/internal/common"
)

type EditState struct {
	handler *StateHandler
}

func NewEditState(rep activity.IActivityRepImpl) (ret StateHandlerImp) {
	return &EditState{NewStateHandler(rep)}
}

func (a EditState) Arraignment(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_ARRAIGNMENT)
	if success {
		ret = "活动提审成功"
	} else {
		ret = "活动状态变更失败"
	}
	return
}

func (a EditState) CheckPass(i int, i2 int) (ret string) {
	return "编辑中不可审核通过"
}

func (a EditState) CheckRefuse(i int, i2 int) (ret string) {
	return "编辑中不可审核拒绝"
}

func (a EditState) CheckRevoke(i int, i2 int) (ret string) {
	return "编辑中不可审核撤销"
}

func (a EditState) Close(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_CLOSE)
	if success {
		ret = "活动关闭成功"
	} else {
		ret = "活动状态变更失败"
	}
	return
}

func (a EditState) Open(i int, i2 int) (ret string) {
	return "非关闭活动不可开启"
}

func (a EditState) Doing(i int, i2 int) (ret string) {
	return "编辑中活动不可执行活动中变更"
}
