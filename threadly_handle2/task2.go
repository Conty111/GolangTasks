/*
Напишите функцию ToString[T any](done <-chan struct{}, valueStream <-chan T)<-chan string,
которая преобразует значения из входного канала в string и записывает в выходной канал.
Используйте fmt.Sprint для преобразования.
Функция должна завершать работу при закрытии канала done, либо при закрытии канала in.
*/

package threadly_handle2

import "fmt"

func ToString[T any](done <-chan struct{}, valueStream <-chan T) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for val := range valueStream {
			select {
			case <-done:
				return
			default:
				out <- fmt.Sprint(val)
			}
		}
	}()
	return out
}
