/*


Релизовать LFU cache LFU — это алгоритм хранения данных в кеше,
который подсчитывает частоту использования каждого элемента и удаляет те,
к которым обращаются реже всего при достижении границ.

Реализуйте эту структуру данных и функции, описанные ниже

// Структура кеша
// Обязательно содержит 2 публичных поля
// UpperBound - верхняя граница размера кеша
// LowerBound - нижняя граница размера кеша
// Если len > UpperBound, кеш автоматически вытеснит значения до нижней границы
// Если любое из этих значений 0 - то этого не произойдет
type Cache struct {
    UpperBound int
    LowerBound int
}


// Создает инстанс кеша
func New() *Cache

// Проверяет, содержит, ли кеш ключ
func (c *Cache) Has(key string) bool

// Возвращает значение по ключу, если оно существует
// Возвращает nil, если не существует
func (c *Cache) Get(key string) interface{} {
    c.lock.Lock()
    defer c.lock.Unlock()
    if e, ok := c.values[key]; ok {
        c.increment(e)
        return e.value
    }
    return nil
}

// Сохраняет значение по ключу
func (c *Cache) Set(key string, value interface{})

// Возвращает размер кеша
func (c *Cache) Len() int

// Возвращает частоту обращений к ключу
func (c *Cache) GetFrequency(key string) int

// Возвращает все ключи в кеше
func (c *Cache) Keys() []string

// Удаляет заданное количество наименее часто используемых элементов элементов
// Возвращает количество удаленных элементов
func (c *Cache) Evict(count int) int


*/

package cash_patterns

import (
	"math"
	"sync"
)

type Cache1 struct {
	UpperBound int
	LowerBound int
	values     map[string]interface{}
	counts     map[string]int
	length     int
	lock       sync.Mutex
}

func New() *Cache1 {
	return &Cache1{
		UpperBound: int(math.Inf(1)),
		LowerBound: 0,
		lock:       sync.Mutex{},
		counts:     make(map[string]int),
		values:     make(map[string]interface{}),
		length:     0,
	}
}

func (c *Cache1) increment(key string) {
	c.counts[key]++
}

func (c *Cache1) Has(key string) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	_, ok := c.values[key]
	return ok
}

func (c *Cache1) Get(key string) interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()
	if e, ok := c.values[key]; ok {
		c.increment(key)
		return e
	}
	return nil
}

func (c *Cache1) Set(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	_, ok := c.values[key]
	if !ok {
		if c.length == c.UpperBound {
			c.Evict(1)
		}
		c.length++
	}
	c.values[key] = value
	c.counts[key] = 1
}

func (c *Cache1) Len() int {
	return c.length
}

func (c *Cache1) GetFrequency(key string) int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.counts[key]
}

func (c *Cache1) Keys() []string {
	c.lock.Lock()
	defer c.lock.Unlock()
	keys := make([]string, c.length)
	var i int
	for key, _ := range c.values {
		keys[i] = key
		i++
	}
	return keys
}

func (c *Cache1) Evict(count int) int {
	//c.lock.Lock()
	//defer c.lock.Unlock()
	lowCount := 1
	deleted := 0
	keys := make([]string, count)
	if count > c.length || c.length-c.LowerBound < count {
		count = c.length - c.LowerBound
	} else if count < 1 {
		return 0
	}
	for deleted < count {
		minBound := 90000000000
		for key, val := range c.counts {
			if val == lowCount {
				keys[deleted] = key
				deleted++
			} else {
				minBound = min(minBound, val)
			}
			if deleted == count {
				break
			}
		}
		if minBound == 90000000000 || deleted == count {
			for _, k := range keys[:deleted] {
				delete(c.values, k)
				delete(c.counts, k)
			}
			c.length -= deleted
		}
		lowCount = minBound
	}
	return deleted
}
