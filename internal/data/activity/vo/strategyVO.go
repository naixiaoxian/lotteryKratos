package vo

import "time"

type StrategyVO struct {
	StrategyId   int64
	StrategyDesc string
	StrategyMode int
	GrantType    int
	GrantDate    time.Time
	ExtInfo      string
	StrategyList []StrategyDetailVO
}
