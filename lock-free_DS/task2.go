package lock_free_DS

import (
	"sync/atomic"
	"unsafe"
)

func AtomicSwap[T any](a *T, b *T) {
	atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(b)),
		atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(a)),
			atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(b)))))
}
