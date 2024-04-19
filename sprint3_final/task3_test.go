/*
Протестируйте следующую реализацию потокобезопасной мапы:

	type SafeMap struct {
	    m   map[string]interface{}
	    mux sync.Mutex
	}

	func NewSafeMap() *SafeMap {
	    return &SafeMap{
	        m: make(map[string]interface{}),
	    }
	}

	func (s *SafeMap) Get(key string) interface{} {
	    s.mux.Lock()
	    defer s.mux.Unlock()

	    return s.m[key]
	}

	func (s *SafeMap) Set(key string, value interface{}) {
	    s.mux.Lock()
	    defer s.mux.Unlock()

	    s.m[key] = value
	}

Примечания
Покрытие кода тестами должно быть 100%.

*/

package sprint3_final

import (
	"testing"
)

func TestSafeMap(t *testing.T) {
	t.Run("Default positive", func(t *testing.T) {
		key := "key"
		value := "value"
		m := NewSafeMap()
		m.Set(key, value)
		if got := m.Get(key); got != value {
			t.Errorf("Get('%v') = '%v', expected: %v", key, value, got)
		}
	})
}
