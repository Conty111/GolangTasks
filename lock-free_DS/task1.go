/*
Счетчик
Напишите структуру-счетчик, который поддерживает инкрементацию и декрементацию с использованием атомиков.

Примечания
Сама структура выглядит таким образом.

type Counter struct {
    value int64
}
Нужно реализовать для нее методы Increment, Decrement, GetValue.
*/

package lock_free_DS

import (
	"sync/atomic"
	"unsafe"
)

type Counter struct {
	value int64
}

func (c *Counter) Increment() {
	atomic.AddInt64((*int64)(unsafe.Pointer(&c.value)), 1)
}

func (c *Counter) Decrement() {
	atomic.AddInt64((*int64)(unsafe.Pointer(&c.value)), -1)
}

func (c *Counter) GetValue() int64 {
	return c.value
}
