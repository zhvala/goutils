package timestamp

import (
	"strconv"
	"time"
)

// Second return millisecond timestamp
func Second() int64 {
	return time.Now().Unix()
}

// SecondStr return millisecond timestamp string
func SecondStr() string {
	return strconv.FormatInt(Second(), 10)
}

// Millisecond return millisecond timestamp
func Millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// MillisecondStr return millisecond timestamp
func MillisecondStr() string {
	return strconv.FormatInt(Millisecond(), 10)
}
