/*
Lock-free очередь
Реализуйте структуру - lock-free очередь,
где операции добавления и извлечения элементов производятся без использования блокировок,
а с использованием CAS (Compare and Swap) на указателях.

Примечания
Сами структуры выглядят таким образом.

type Node struct {
    value int
    next  *Node
}

type LockFreeQueue struct {
    head unsafe.Pointer
    tail unsafe.Pointer
}
Нужно реализовать для них методы.

NewLockFreeQueue() *LockFreeQueue
Enqueue(value int)
Dequeue() (int, bool)
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

type LockFreeQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewLockFreeQueue() *LockFreeQueue {
	return &LockFreeQueue{}
}

func (q *LockFreeQueue) Enqueue(value int) {
	if q.head != nil {
		node := &Node{value: value}
		oldTail := (*Node)(atomic.LoadPointer(&q.tail))
		node.next = oldTail
		atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(oldTail), unsafe.Pointer(node))
	} else {
		q.head = unsafe.Pointer(&Node{value: value})
	}
}

func (q *LockFreeQueue) Dequeue() (int, bool) {
	var val int
	var res bool
	if q.head != nil {
		oldHead := (*Node)(atomic.LoadPointer(&q.head))
		val = oldHead.value
		res = true
		if q.tail == nil {
			q.head = nil
		}
	}
	if q.tail != nil {
		var prev unsafe.Pointer
		currPtr := q.tail
		curr := (*Node)(currPtr)
		for curr.next != nil {
			prev = currPtr
			atomic.CompareAndSwapPointer(&currPtr, currPtr, unsafe.Pointer(curr.next))
			curr = (*Node)(currPtr)
		}
		node := (*Node)(prev)
		node.next = nil
		if q.head == nil {
			val = curr.value
			res = true
		}
		atomic.CompareAndSwapPointer(&q.head, q.head, prev)
	}
	return val, res
}
