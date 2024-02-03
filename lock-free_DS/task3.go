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

type Node1 struct {
	value int
	next  *Node1
}

type LockFreeQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewLockFreeQueue() *LockFreeQueue {
	node := &Node1{}
	return &LockFreeQueue{
		head: unsafe.Pointer(node),
		tail: unsafe.Pointer(node),
	}
}

func (q *LockFreeQueue) Enqueue(value int) {
	newNode := &Node1{value: value}

	for {
		tail := (*Node1)(atomic.LoadPointer(&q.tail))
		next := tail.next

		if next == nil {
			// Попытка добавить новый узел
			if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&tail.next)), nil, unsafe.Pointer(newNode)) {
				// Успешно добавлен новый узел, обновляем tail
				atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(tail), unsafe.Pointer(newNode))
				return
			}
		} else {
			// Узел уже добавлен другой горутиной, обновляем tail
			atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(tail), unsafe.Pointer(next))
		}

	}
}

func (q *LockFreeQueue) Dequeue() (int, bool) {
	for {
		head := q.head
		tail := q.tail
		first := unsafe.Pointer((*Node1)(atomic.LoadPointer(&head)).next)

		if head == q.head {
			if head == tail {
				// Очередь пуста
				if first == nil {
					return 0, false
				}
				// Обновляем tail, так как другие горутины добавили узлы
				atomic.CompareAndSwapPointer(&q.tail, tail, first)
			} else {
				// Читаем значение из узла
				value := (*Node1)(first).value
				// Пробуем обновить head
				if atomic.CompareAndSwapPointer(&q.head, head, first) {
					return value, true
				}
			}
		}
	}
}
