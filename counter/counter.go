package counter

import (
	"fmt"
	"sync/atomic"
)

// Just Counter, Simple Enough, Easy to Incr/Decr
// c := Counter(0)
// c.SetCount(100)
// c.Incr(1) or c.Decr(1)
// fmt.Println(c) or c.Count()
type Counter int64

func (c *Counter) SetCount(val int64) {
	atomic.StoreInt64((*int64)(c), val)
}

func (c *Counter) Count() int64 {
	return atomic.LoadInt64((*int64)(c))
}

func (c *Counter) Incr(delta int64) int64 {
	return atomic.AddInt64((*int64)(c), delta)
}

func (c *Counter) Decr(delta int64) int64 {
	return atomic.AddInt64((*int64)(c), delta*-1)
}

func (c *Counter) String() string {
	return fmt.Sprint(c.Count())
}
