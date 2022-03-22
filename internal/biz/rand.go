package biz

import (
	"github.com/shopspring/decimal"
	"github.com/valyala/fastrand"
	"math/rand"
	"time"
)

var (
	rngFast fastrand.RNG
	rngMath = rand.New(rand.NewSource(time.Now().Unix()))
)

func RInt(n int) int {
	return int(rngFast.Uint32n(uint32(n)))
}

func RDecimal(val, exp int) decimal.Decimal {
	in := rngMath.Float64() + float64(RInt(val))
	return decimal.NewFromFloatWithExponent(in, int32(exp))
}
