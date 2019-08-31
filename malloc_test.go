package malloc_test

import (
	"testing"
	"unsafe"

	"github.com/skillian/malloc"
)

func TestHugeAlloc(t *testing.T) {
	var allocs []unsafe.Pointer
	for {
		p, length := getHugeAlloc()
		if p == nil {
			break
		}
		allocs = append(allocs, p)
		t.Log("allocated:", length)
	}
	for _, a := range allocs {
		malloc.Free(a)
	}
}

func getHugeAlloc() (unsafe.Pointer, uintptr) {
	length := uintptr(1 << 30)
	for length > (1 << 24) {
		p := malloc.Malloc(length)
		if p != nil {
			return p, length
		}
		length >>= 1
	}
	return nil, 0
}
