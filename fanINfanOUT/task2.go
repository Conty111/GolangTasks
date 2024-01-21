/*
Синхронизация Генератор чисел
Дан интерфейс:

type sequenced interface {
	getSequence() int
}

Напишите функцию func EvenNumbersGen[T sequenced](ctx context.Context, numbers ...T) <-chan T,
которая получает на вход элементы типа T, а возвращает канал,
который выдает только те элементы из списка, для которых getSequence возвращает четное число.
Код решения должен содержать объявление интерфейса sequenced.
*/

package fanINfanOUT

import (
	"context"
)

type Sequenced interface {
	GetSequence() int
}

func EvenNumbersGen[T sequenced](ctx context.Context, numbers ...T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for _, num := range numbers {
			if num.GetSequence()%2 == 0 {
				out <- num
			}
		}
	}()
	return out
}
