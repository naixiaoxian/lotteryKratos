package req

import "time"

type PartakeReq struct {
	Uid         string
	ActivityId  int32
	PartakeData time.Time
}
