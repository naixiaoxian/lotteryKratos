package vo

import "time"

type StrategyBriefVO struct {
	StrategyId   int64
	StrategyDesc string
	StrategyMode int
	GrantType    int
	GrandDate    time.Time
	ExtInfo      string
}
