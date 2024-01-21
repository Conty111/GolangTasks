/*
В некоторых случаях на одном из этапов обработки данных нужно сделать копию входящего потока.
Например, для записи лога. Напишите функцию Tee[T any](done <-chan struct{}, in <-chan T) (<-chan T, <-chan T),
которая записывает данные из канала in в два выходящих канала.
Функция должна завершать работу при закрытии канала done, либо при закрытии канала in.
*/

package threadly_handle2

func Tee[T any](done <-chan struct{}, in <-chan T) (<-chan T, <-chan T) {
	out := make(chan T)
	outLog := make(chan T)
	go func() {
		defer close(out)
		defer close(outLog)
		for val := range in {
			select {
			case <-done:
				return
			default:
				out <- val
				outLog <- val
			}
		}
	}()
	return out, outLog
}
