package main

import (
	// "github.com/latermoon/libs/counter"
	"../counter"
	"fmt"
)

func main() {
	c := counter.Counter(0)
	c.SetCount(10)
	fmt.Println("count:", c.Incr(1), c.Decr(2))

	factory := counter.NewFactory()
	c2 := factory.Get("name")
	c2.Incr(1)
	factory.Get("name").Incr(2)
	fmt.Println(factory.Get("name").Decr(10))
}
