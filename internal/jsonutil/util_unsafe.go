package jsonutil

import (
	"reflect"
	"unsafe"
)

// strToBytes converts s into slice of bytes with out copying it.
// The resulting slice must not be stored or modified!
func strToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data,
		Len:  len(s),
		Cap:  len(s),
	}))
}
