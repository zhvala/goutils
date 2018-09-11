package strtools

import (
	"reflect"
	"unsafe"
)

// BytesToString is a fast method which convert byte slice
// to string. Any changes of []byte or string will change
// the other one. Since this function does not create
// a new string, it simply changes the way it is interpreted.
func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
