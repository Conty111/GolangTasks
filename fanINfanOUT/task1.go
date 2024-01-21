package fanINfanOUT

import (
	"context"
	"sync"
)

// FanIn - объединяет данные из нескольких каналов в один.
func FanIn[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	outputCh := make(chan T)      // выходной канал
	wg := sync.WaitGroup{}        // для ожидания завершения
	for _, ch := range channels { // переберём все каналы
		wg.Add(1)
		// для каждого канала вызовем функцию в отдельной горутине
		go func(input <-chan T) {
			defer wg.Done() // отметим, что функция завершилась
			for {           // цикл для получения данных из входных каналов
				select {
				case data, ok := <-input: // получим данные из канала
					if !ok {
						return // данных больше нет - выходим
					}
					outputCh <- data // запишем данные в выходной канал
				case <-ctx.Done(): // если нужно завершить - выходим
					return
				}
			}
		}(ch) // передадим входной канал в функцию
	}
	go func() {
		wg.Wait()       // дождёмся завершения обработки всех каналов
		close(outputCh) // закроем выходной канал
	}()

	return outputCh // вернём канал
}
