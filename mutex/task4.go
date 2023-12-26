/*
Напишите функцию Write(num int), которая записывает данный в буфер Buf []int.
Функция Consume() int должна забирать первое значение из этого буфера и возвращать его.
Используйте мьютекс для синхронизации доступа к буферу.
*/

package mutex

import "sync"

var Buf []int
var mu sync.Mutex = sync.Mutex{}

func Write(num int) {
	mu.Lock()
	Buf = append(Buf, num)
	mu.Unlock()
}

func Consume() int {
	mu.Lock()
	defer mu.Unlock()
	var val int
	if Buf != nil {
		val = Buf[0]
		Buf = Buf[1:]
	}
	return val
}
