package memory

import (
	"reflect"
	"unsafe"
)

// View the memory of data.
func View(data any) []byte {
	v := reflect.ValueOf(data).Elem()
	ptr := unsafe.Pointer(uintptr(v.Addr().UnsafePointer()))
	size := v.Type().Size()

	return (*[1 << 30]byte)(ptr)[:size:size]
}
