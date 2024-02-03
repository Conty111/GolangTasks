/*
Lock-free стэк *



Реализуйте структуру - lock-free стэк,
где операции добавления и извлечения элементов производятся без использования блокировок.
Примечания

Сами структуры выглядят таким образом.

type Node struct {
    value int
    next  *Node
}

type LockFreeStack struct {
    top unsafe.Pointer
}

Нужно реализовать для них методы.

NewLockFreeStack() *LockFreeStack
Push(value int)
Pop() (int, bool)

*/

package lock_free_DS

import (
	"sync/atomic"
	"unsafe"
)

type Node struct {
	value int
	next  *Node
}

type LockFreeStack struct {
	top unsafe.Pointer
}

func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{
		top: unsafe.Pointer(&Node{}),
	}
}

func (st *LockFreeStack) Push(value int) {
	topPtr := atomic.LoadPointer((*unsafe.Pointer)(st.top))
	top := (*Node)(topPtr)
	newNode := &Node{value: value, next: top}
	atomic.CompareAndSwapPointer((*unsafe.Pointer)(st.top), topPtr, unsafe.Pointer(newNode))
}

func (st *LockFreeStack) Pop() (int, bool) {
	topPtr := atomic.LoadPointer((*unsafe.Pointer)(st.top))
	top := (*Node)(topPtr)
	if top == nil {
		return 0, false
	}
	atomic.CompareAndSwapPointer((*unsafe.Pointer)(st.top), topPtr, unsafe.Pointer(top.next))
	return top.value, true
}
