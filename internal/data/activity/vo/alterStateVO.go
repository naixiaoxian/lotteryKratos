package vo

// AlterStateVO 变更活动状态对象
type AlterStateVO struct {
	ActivityId  int64
	BeforeState int
	AfterState  int
}
