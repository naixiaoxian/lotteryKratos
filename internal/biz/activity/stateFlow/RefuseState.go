package stateFlow

import (
	"lotteryKratos/internal/biz/activity"
	"lotteryKratos/internal/common"
)

type RefuseState struct {
	handler *StateHandler
}

func NewRefuseState(rep activity.IActivityRepImpl) (ret StateHandlerImp) {
	return &RefuseState{NewStateHandler(rep)}
}

func (a RefuseState) Arraignment(i int, i2 int) (ret string) {
	return "已审核不可重复提审"
}

func (a RefuseState) CheckPass(i int, i2 int) (ret string) {
	return "已审核状态不可重复审核"
}

func (a RefuseState) CheckRefuse(i int, i2 int) (ret string) {
	return "活动审核拒绝不可重复审核"
}

func (a RefuseState) CheckRevoke(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_REVOKE)
	if success {
		ret = "活动审核拒绝完成"
	} else {
		ret = "活动状态变更失败"
	}
	return
}

func (a RefuseState) Close(i int, i2 int) (ret string) {
	success := a.handler.activityrepImpl.AlterStatus(int64(i), i2, common.ACTIVITY_STATE_CLOSE)
	if success {
		ret = "活动审核关闭完成"
	} else {
		ret = "活动状态变更失败"
	}
	return
}

func (a RefuseState) Open(i int, i2 int) (ret string) {
	return "非关闭活动不可开启"
}

func (a RefuseState) Doing(i int, i2 int) (ret string) {
	return "审核拒绝不可执行活动为进行中"
}
