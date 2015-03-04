package protocol

import (
	"strconv"
)

const (
	CR   = '\r'
	LF   = '\n'
	CRLF = "\r\n"
)

var itoaCache []string

func init() {
	itoaCache = make([]string, 1000)
	for i, count := 0, len(itoaCache); i < count; i++ {
		itoaCache[i] = strconv.Itoa(i)
	}
}

// itoa speed up the strconv.Itoa in small numbers
func itoa(i int) string {
	if i > 0 && i < len(itoaCache) {
		return itoaCache[i]
	} else {
		return strconv.Itoa(i)
	}
}
