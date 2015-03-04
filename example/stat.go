package main

import (
	"github.com/latermoon/libs/stat"
	"math/rand"
	"os"
)

//     time   count  status
// 15:18:51     581    WAIT
// 15:18:52     887    WAIT
// 15:18:53     847    WAIT
// 15:18:54     559    WAIT
// 15:18:55     581    WAIT
func main() {
	w := stat.New(os.Stdout)
	w.Add(stat.TextItem("time", 8, func() interface{} {
		return stat.TimeString()
	}))
	w.Add(stat.TextItem("count", 8, func() interface{} {
		return 500 + rand.Intn(500)
	}))
	w.Add(stat.TextItem("status", 8, func() interface{} {
		return "WAIT"
	}))
	w.Start()
}
