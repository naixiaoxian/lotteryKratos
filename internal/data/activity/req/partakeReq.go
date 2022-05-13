package req

import "time"

type PartakeReq struct {
	Uid         string
	ActivityId  int64
	PartakeDate time.Time
}
