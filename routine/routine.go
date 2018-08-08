package routine

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// GoroutineID get current goroutine id
// this method is slow
// should not call it frequently
func GoroutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
