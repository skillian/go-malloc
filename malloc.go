package malloc

/*
#include <stdlib.h>

void*
github_com_skillian_malloc_malloc(size_t size) {
	return malloc(size);
}

void*
github_com_skillian_malloc_realloc(void* ptr, size_t size) {
	return realloc(ptr, size);
}
*/
import "C"
import "unsafe"

// Malloc calls malloc but does not crash Go if the allocation fails.
func Malloc(size uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.github_com_skillian_malloc_malloc(C.size_t(size)))
}

// Realloc calls realloc with the given pointer and size.
func Realloc(pointer unsafe.Pointer, size uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.github_com_skillian_malloc_realloc(pointer, C.size_t(size)))
}

// Free the pointer.
func Free(pointer unsafe.Pointer) {
	C.free(pointer)
}
