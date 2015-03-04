package counter

import (
	"sync"
)

// Counter Factory & Container
type Counters struct {
	table map[string]Counter
	mu    sync.Mutex
}

func NewCounters() (c *Counters) {
	c = &Counters{
		table: make(map[string]Counter),
	}
	return
}

func (c *Counters) Len() int {
	return len(c.table)
}

// 获取并自动创建
func (c *Counters) Get(name string) (counter Counter) {
	var ok bool
	counter, ok = c.table[name]
	if !ok {
		c.mu.Lock()
		counter, ok = c.table[name]
		if !ok {
			counter = New(0)
			c.table[name] = counter
		}
		c.mu.Unlock()
	}
	return counter
}

func (c *Counters) Names() (names []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	names = make([]string, 0, len(c.table))
	for key, _ := range c.table {
		names = append(names, key)
	}
	return
}
