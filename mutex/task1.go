/*
Реализуйте потокобезопасную мапу.
Для чтения элементов используйте функцию func (s *SafeMap) Get(key string) interface{},
а для записи func (s *SafeMap) Set(key string, value interface{}) .
Используйте func NewSafeMap() *SafeMap для получению нового экземпляра.

Примечания
Код должен содержать структуру:

type SafeMap struct {
m map[string]interface{}
mux sync.Mutex
}
*/

package mutex

import "sync"

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m:   make(map[string]interface{}),
		mux: sync.Mutex{},
	}
}

func (s *SafeMap) Get(key string) interface{} {
	s.mux.Lock()
	val := s.m[key]
	s.mux.Unlock()
	return val
}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m[key] = value
}
