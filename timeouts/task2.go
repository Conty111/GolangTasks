/*
Сгенерируйте N-е число Фибоначчи.
Установите ограничение в T и, если оно пройдет в течение указанного времени,
выведите число Фибоначчи в одной строке без кавычек “Fibonacci(N) = D”,
где D - N-е число Фибоначчи, в противном случае выдайте ошибку.
Сигнатура функции TimeoutFibonacci(n int, timeout time.Duration) (int, error)
n - номер числа timeout - время, отведенное на операцию
*/

package timeouts

import (
	"fmt"
	"time"
)

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	ch := make(chan int)
	go func() {
		a, b := 0, 1
		if n < 2 {
			ch <- n
			return
		}
		for i := 0; i < n-1; i++ {
			a, b = a+b, a
		}
		ch <- a
	}()
	select {
	case res := <-ch:
		if res < 0 {
			return res, fmt.Errorf("Invalid n")
		}
		return res, nil
	case <-time.After(timeout):
		return 0, fmt.Errorf("Timeout error")
	}
}
