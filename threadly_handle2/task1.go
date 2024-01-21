/*
Напишите функцию DoubleNumbers(done <-chan struct{}, in <-chan int) <-chan int,
чтобы она удваивала элементы из канала in и записывала их в выходной канал.
Функция должна завершать работу при закрытии канала done, либо при закрытии канала in.
*/

package threadly_handle2

func DoubleNumbers(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for val := range in {
			select {
			case <-done:
				return
			default:
				out <- val * 2
			}
		}
	}()
	return out
}
