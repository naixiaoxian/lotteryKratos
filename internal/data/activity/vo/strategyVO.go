package vo

import "time"

type StrategyVO struct {
	StrategyId   string
	StrategyDesc string
	StrategyMode int
	GrantType    int
	GrantDate    time.Time
	ExtInfo      string
	StrategyList []StrategyDetailVO
}
