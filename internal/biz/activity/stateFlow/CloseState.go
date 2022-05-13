package stateFlow

import (
	"lotteryKratos/internal/biz/activity"
	"lotteryKratos/internal/common"
)

type CloseState struct {
	handler *StateHandler
}

func NewCloseState(rep activity.IActivityRepImpl) (ret StateHandlerImp) {
	return &CloseState{NewStateHandler(rep)}
}

func (a CloseState) Arraignment(i int, i2 int) (ret string) {
	return "活动关闭不可提审"
}

func (a CloseState) CheckPass(i int, i2 int) (ret string) {
	return "活动关闭不可审核通过"
}

func (a CloseState) CheckRefuse(i int, i2 int) (ret string) {
	return "活动关闭不可拒绝"
}

func (a CloseState) CheckRevoke(i int, i2 int) (ret string) {
	return "活动关闭不可撤销"
}

func (a CloseState) Close(i int, i2 int) (ret string) {
	return "活动关闭不可重复关闭"
}

func (a CloseState) Open(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_OPEN)
	if success {
		ret = "活动开启完成"
	} else {
		ret = "活动状态变更失败"
	}
	return
}

func (a CloseState) Doing(i int, i2 int) (ret string) {
	return "活动关闭不可变更活动中"
}
