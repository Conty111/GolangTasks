/*
Напишите программу для генерации простых чисел вплоть до N (2 < N < 105)
и используйте для реализации паттерна таймаута.
Выводите каждое простое число на новой строке.
Через 0,01 секунду и остановите генерацию простых чисел.

Реализуйте функцию
func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int)
stop - канал для остановки генерации
prime_nums - канал для вывода простых чисел
N - число до которого нужно генерировать числа
*/

package timeouts_pattern

import "time"

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	N++
	arr := make([]bool, N)
	defer close(prime_nums)
	defer close(stop)
	timeout := time.After(100 * time.Millisecond)
	for i := 2; i < N-2; i++ {
		select {
		case <-timeout:
			return
		case <-stop:
			return
		default:
			if !arr[i-2] {
				for j := i + 1; j < N-2; j += i {
					arr[j-2] = true
				}
				prime_nums <- i
			}
		}
	}
}
