package counter

import (
	"fmt"
	"sync/atomic"
)

// Just Counter, Simple Enough, Easy to Incr/Decr
type Counter struct {
	v int64
}

func New(val int64) *Counter {
	return &Counter{
		v: val,
	}
}

func (c *Counter) SetCount(val int64) {
	atomic.StoreInt64(&c.v, val)
}

func (c *Counter) Count() int64 {
	return atomic.LoadInt64(&c.v)
}

func (c *Counter) Incr(delta int64) int64 {
	return atomic.AddInt64(&c.v, delta)
}

func (c *Counter) Decr(delta int64) int64 {
	return atomic.AddInt64(&c.v, delta*-1)
}

func (c *Counter) String() string {
	return fmt.Sprint(c.Count())
}
