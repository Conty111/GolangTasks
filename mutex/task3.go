/*
Напишите потокобезопасную очередь ConcurrentQueue. Реализуете следующий интерфейс:

type Queue interface {
Enqueue(element interface{}) // положить элемент в очередь
Dequeue() interface{} // забрать первый элемент из очереди
}

Примечания
Код должен содержать следующую структуру:

type ConcurrentQueue struct {
queue []interface{} // здесь хранить элементы очереди
mutex sync.Mutex
}
*/

package mutex

import "sync"

type Queue interface {
	Enqueue(element interface{}) // положить элемент в очередь
	Dequeue() interface{}        // забрать первый элемент из очереди
}

type ConcurrentQueue struct {
	queue []interface{} // здесь хранить элементы очереди
	mutex sync.Mutex
}

func (q *ConcurrentQueue) Enqueue(element interface{}) {
	q.mutex.Lock()
	q.queue = append(q.queue, element)
	q.mutex.Unlock()
}

func (q *ConcurrentQueue) Dequeue() interface{} {
	if q.queue != nil {
		val := q.queue[0]
		q.mutex.Lock()
		q.queue = q.queue[1:]
		q.mutex.Unlock()
		return val
	}
	return nil
}
