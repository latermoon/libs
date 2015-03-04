package counter

import (
	"sync/atomic"
)

// Just Counter, Simple Enough, Easy to Incr/Decr
type Counter int64

func New(val int64) Counter {
	return Counter(val)
}

func (c Counter) Incr(delta int64) int64 {
	return atomic.AddInt64(&c, delta)
}

func (c Counter) Decr(delta int64) int64 {
	return atomic.AddInt64(&c, delta*-1)
}

func (c Counter) SetCount(val int64) {
	atomic.StoreInt64(&c, val)
}

func (c Counter) Count() int64 {
	return atomic.LoadInt64(&c)
}
