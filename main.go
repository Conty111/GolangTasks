package main

import (
	"context"
	"fmt"
	"github.com/Conty111/GolangTasks/fanINfanOUT"
	"sync"
)

type Num struct {
	data int
}

func (t *Num) GetSequence() int {
	return t.data
}

func main() {
	nums := []*Num{
		{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9},
	}
	nums1 := []*Num{
		{10}, {11}, {12}, {13}, {14}, {15}, {16},
	}
	ctx := context.Background()
	inputCh1 := fanINfanOUT.EvenNumbersGen(ctx, nums...)
	inputCh2 := fanINfanOUT.EvenNumbersGen(ctx, nums1...)
	//inputCh2 := fanINfanOUT.OddNumbersGen(ctx, nums...)
	// запись во временную очередь
	inCh := inTemp(ctx, inputCh1, inputCh2)
	// обработка из временной очереди и упорядочивание
	outCh := fanINfanOUT.ProcessTempCh(ctx, 2, inCh)

	for num := range outCh {
		fmt.Println(num) // {0} {1} {2} {3} {4} {5} {6}
	}
}

func inTemp[T fanINfanOUT.Sequenced](
	ctx context.Context,
	channels ...<-chan T,
) <-chan fanINfanOUT.FanInRecord[T] {
	// канал для ожидания
	fanInCh := make(chan fanINfanOUT.FanInRecord[T])
	// для синхронизации
	wg := sync.WaitGroup{}
	// перебор всех входных каналов
	for i := range channels {
		wg.Add(1)
		// запустим горутину для получения данных из канала
		go func(index int) {
			defer wg.Done()
			// канал для синхронизации
			pauseCh := make(chan struct{})
			// цикл для получения данных из канала
			for {
				select {
				// получим данные из канала
				case data, ok := <-channels[index]:
					if !ok {
						return // канал закрыт - выходим
					}
					// положим во временный канал вместе с индексом
					fanInCh <- fanINfanOUT.FanInRecord[T]{
						// индекс канала, откуда пришли данные
						Index: index,
						// данные из канала
						Data: data,
						// канал для синхронизации
						Pause: pauseCh,
					}
				case <-ctx.Done():
					return
				}
				// ждём, пока в канал pause не будет передан сигнал
				// о получении очередного элемента из канала
				select {
				case <-pauseCh:
				// сняли с паузы
				// продолжим обработку данных из входного канала
				case <-ctx.Done():
					return
				}
			}
		}(i)
	}
	go func() {
		// ожидаем завершения
		wg.Wait()
		close(fanInCh)
	}()
	// вернём канал с неотсортированными элементами
	return fanInCh
}
