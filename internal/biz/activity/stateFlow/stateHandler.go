package stateFlow

import (
	"lotteryKratos/internal/biz/activity"
	"lotteryKratos/internal/common"
)

type StateHandlerImp interface {
	Arraignment(int, int) string
	CheckPass(int, int) string
	CheckRefuse(int, int) string
	CheckRevoke(int, int) string
	Close(int, int) string
	Open(int, int) string
	Doing(int, int) string
}

type StateHandler struct {
	activityrepImpl activity.IActivityRepImpl
}

func NewStateHandler(rep activity.IActivityRepImpl) (ret *StateHandler) {
	return &StateHandler{activityrepImpl: rep}
}

var StateGroups = make(map[int]StateHandlerImp)

func InitGroups(rep activity.IActivityRepImpl) {
	StateGroups[common.ACTIVITY_STATE_ARRAIGNMENT] = NewArraignmentState(rep)
	StateGroups[common.ACTIVITY_STATE_CLOSE] = NewCloseState(rep)
	StateGroups[common.ACTIVITY_STATE_DOING] = NewDoingState(rep)
	StateGroups[common.ACTIVITY_STATE_EDIT] = NewEditState(rep)
	StateGroups[common.ACTIVITY_STATE_OPEN] = NewOpenState(rep)
	StateGroups[common.ACTIVITY_STATE_PASS] = NewPassState(rep)
	StateGroups[common.ACTIVITY_STATE_REFUSE] = NewRefuseState(rep)

}

func GetHandler(states int) StateHandlerImp {
	return StateGroups[states]
}
