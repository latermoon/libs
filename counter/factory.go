package counter

import (
	"sync"
)

// Counter Factory & Collection
// factory := NewFactory()
// factory.Get("set").Incr(1)
// factory.Get("get").Incr(1)
// factory.Get("del").Incr(1)
// factory.Get("total").Incr(3)
type Factory struct {
	table map[string]*Counter
	mu    sync.Mutex
}

func NewFactory() (f *Factory) {
	f = &Factory{
		table: make(map[string]*Counter),
	}
	return
}

// Get or auto create a Counter by name
func (f *Factory) Get(name string) (c *Counter) {
	var ok bool
	if c, ok = f.table[name]; !ok {
		f.mu.Lock()
		if c, ok = f.table[name]; !ok {
			tmp := Counter(0)
			c = &tmp
			f.table[name] = c
		}
		f.mu.Unlock()
	}
	return c
}

func (f *Factory) Len() int {
	return len(f.table)
}
