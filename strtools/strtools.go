package strtools

import (
	"reflect"
	"unsafe"
)

// BytesToString is a fast method which convert byte slice
// to string. Any changes of []byte or string will change
// the other one. Since this function does not create
// a new string, it simply changes the way it is interpreted.
// Notice, this method is not safe.
func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

// StringToBytes is a fast method which convert string to
// byte slice. Any changes of []byte or string will change
// the other one. Since this function does not create
// a new []byte, it simply changes the way it is interpreted.
// Notice, this method is not safe.
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
