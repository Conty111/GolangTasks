package lock_free_DS

import (
	"sync/atomic"
	"unsafe"
)

func AtomicSwap[T any](p1 *any, p2 *any) {
	res := unsafe.Pointer(p1)
	atomic.CompareAndSwapPointer(&res, unsafe.Pointer(p1), unsafe.Pointer(p2))
}
