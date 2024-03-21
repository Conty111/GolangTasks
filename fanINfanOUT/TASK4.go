/*
Синхронизация OrderedFanIn
На вход программы поступают сообщения в виде объектов, тип которых удовлетворяет интерфейсу:

type ordered interface {
	getIndex() int
	getData() string
}

Напишите функцию OrderedFanIn[T ordered](ctx context.Context, channels ...<-chan T) <-chan T,
которая запишет данные в выходной канал в соответствии со значением getIndex()(в порядке увеличения).
Код программы должен содержать определение интерфейса.
*/

package fanINfanOUT

import "context"

type ordered interface {
	getIndex() int
	getData() string
}

func OrderedFanIn[T ordered](ctx context.Context, channels ...<-chan T) <-chan T {
	out := make(chan T)
	go func(outputCh chan T) {
		defer close(outputCh)
		// порядковый номер очередного элемента
		expected := 0
		// буфер для ожидания элементов по количеству входных каналов
		queuedData := make(map[int]T)
		for _, in := range channels {
			// если получили элемент с номером, который ожидаем
			select {
			// запишем элемент в выходной канал
			case inData, ok := <-in:
				if !ok {
					continue
				}
				if inData.getIndex() == expected {
					out <- inData
					expected++
					data, ok := queuedData[expected]
					for ok {
						out <- data
						delete(queuedData, expected)
						expected++
						data, ok = queuedData[expected]
					}
				} else {
					queuedData[inData.getIndex()] = inData
				}
			case <-ctx.Done():
				return
			}
		}
	}(out)
	return out
}
