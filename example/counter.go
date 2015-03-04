package main

import (
	// "github.com/latermoon/libs/counter"
	"../counter"
	"fmt"
)

func main() {
	c := counter.New(0)
	c.Incr(100)
	fmt.Println(c)
}
