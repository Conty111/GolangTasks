/*


MRU (most recently used) cache.
В случае если при установке элемента достигнуто максимальное значение размера кеша,
то удаляем из кеша последний использованный элемент.

Название структуры и сигнатура функций определены ниже.

// структура MRU кеша
type MRUCache struct {}

// возвращает новый инстанс кеша размером capacity
func NewMRUCache(capacity int) *MRUCache
// устанавливает значени value ключу key
func (c *MRUCache) Set(key, value string)
// получает значение и флаг его начличия по ключу key
func (c *MRUCache) Get(key string) (string, bool)


*/

package cash_patterns

import (
	"container/list"
	"sync"
)

// создание LruCache, представляющего кеш LRU
type MRUCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
	mutex    sync.Mutex
}

// создание записи кеша, представляющей запись в кеше LRU
type CacheEntry struct {
	key   string
	value string
}

// NewLRUCache создает новый экземпляр LruCache с указанной ёмкостью
func NewMRUCache(capacity int) *MRUCache {
	return &MRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

// извлечение значения, связанного с данным ключом, из кеша
func (lru *MRUCache) Get(key string) (string, bool) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	if element, ok := lru.cache[key]; ok {
		// перемещение доступного элемента в начало списка (последний раз использованный)
		lru.list.MoveToFront(element)
		return element.Value.(*CacheEntry).value, true
	}
	return "", false
}

// добавление или обновление пары ключ-значение в кеше
func (lru *MRUCache) Set(key, value string) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	// проверка того, существует ли ключ уже в кэше
	if element, ok := lru.cache[key]; ok {
		// обновляем существующую запись и перемещаем ее в начало (использовалась в последний раз)
		element.Value.(*CacheEntry).value = value
		lru.list.MoveToFront(element)
	} else {
		// добавление новой записи в кеш
		entry := &CacheEntry{key: key, value: value}
		element := lru.list.PushFront(entry)
		lru.cache[key] = element

		// проверяем, заполнен ли кеш (есть ли место), при необходимости удаляем наименее недавно использованный элемент
		if lru.list.Len() > lru.capacity {
			newest := lru.list.Back()
			if newest != nil {
				delete(lru.cache, newest.Value.(*CacheEntry).key)
				lru.list.Remove(newest)
			}
		}
	}
}

// PrintCache, который печатает текущее содержимое кеша
//func (lru *LRUCache) PrintCache() {
//	lru.mutex.Lock()
//	defer lru.mutex.Unlock()
//
//	fmt.Printf("LRU Cache (Capacity: %d, Size: %d): [", lru.capacity, lru.list.Len())
//	for element := lru.list.Front(); element != nil; element = element.Next() {
//		entry := element.Value.(*CacheEntry)
//		fmt.Printf("(%s: %v) ", entry.key, entry.value)
//	}
//	fmt.Println("]")
//}
