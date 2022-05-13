package stateFlow

import (
	"lotteryKratos/internal/biz/activity"
	"lotteryKratos/internal/common"
)

type OpenState struct {
	handler *StateHandler
}

func NewOpenState(rep activity.IActivityRepImpl) (ret StateHandlerImp) {
	return &EditState{NewStateHandler(rep)}
}

func (a OpenState) Arraignment(i int, i2 int) (ret string) {
	return "活动开启不可提交审核"
}

func (a OpenState) CheckPass(i int, i2 int) (ret string) {
	return "活动开启不可通过审核"
}

func (a OpenState) CheckRefuse(i int, i2 int) (ret string) {
	return "活动开启不可拒绝审核"
}

func (a OpenState) CheckRevoke(i int, i2 int) (ret string) {
	return "活动开启不可撤销审核"
}

func (a OpenState) Close(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_CLOSE)
	if success {
		ret = "活动关闭成功"
	} else {
		ret = "活动状态变更失败"
	}
	return
}

func (a OpenState) Open(i int, i2 int) (ret string) {
	return "活动开启不可重复开启"
}

func (a OpenState) Doing(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_DOING)
	if success {
		ret = "活动开启完成"
	} else {
		ret = "活动状态变更失败"
	}
	return
}
