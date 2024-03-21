/*
Напишите потокобезопасный счётчик Counter. Реализуете следующий интерфейс:

type Сount interface{
Increment() // увеличение счётчика на единицу
GetValue() int // получение текущего значения
}

Примечания
Код должен содержать следующую структуру:

type Counter struct {
value int // значение счетчика
mu sync.RWMutex
}
*/

package mutex

import "sync"

type Сount interface {
	Increment()    // увеличение счётчика на единицу
	GetValue() int // получение текущего значения
}

type Counter struct {
	value int // значение счетчика
	mu    sync.RWMutex
}

func (c *Counter) Increment() {
	c.mu.RLock()
	c.value++
	c.mu.RUnlock()
}

func (c *Counter) GetValue() int {
	return c.value
}
