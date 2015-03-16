package main

import (
	"../bitmaps"
	"fmt"
)

func main() {
	b := bitmaps.New(30)
	b.SetBit(2, true)
	b.SetBit(6, true)
	b.SetBit(29, true)
	b.SetBit(5, false)
	b.SetBit(6, false)
	for i := 0; i < 30; i++ {
		exist := b.GetBit(i)
		if exist {
			fmt.Println(i, exist)
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("size:", b.Size())
}
