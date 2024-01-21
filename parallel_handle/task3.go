/*
Параллельная обработка Summer
Напишите программу для подсчета суммы чисел в слайсе параллельно, используя горутины.
Программа должна разделять исходный слайс на несколько слайсов и
затем находить сумму каждой части в отдельной горутине. После этого подсчитывается финальная сумма.

Программа должна реализовать следующий интерфейс:

type Summer interface{
	// функция для нахождения суммы чисел
	func ProcessSum(
		// функция, которая будет вызываться для нахождения суммы части слайса.
		// результат суммы записать в result
		summer func(arr []int, result chan<- int),
		// слайс с числами, сумму которых нужно найти
		nums []int,
		// сколько элементов в одной части (последняя часть может быть меньше)
		сhunkSize int,
	) (int, error) // вернуть сумму чисел
}
Также нужно реализовать функцию SumChunk(arr []int, result chan<- int),
которая будет вызываться для нахождения суммы части слайса (summer).
В случае возникновения ошибок верните 0 и возникшую ошибку.
*/

package gzipper

import (
	"fmt"
	"sync"
)

type Summer interface {
	// функция для нахождения суммы чисел
	ProcessSum(summer func(arr []int, result chan<- int), nums []int, сhunkSize int) (int, error)
}

func ProcessSum(
	summer func(arr []int, result chan<- int),
	nums []int,
	сhunkSize int) (int, error) {
	if сhunkSize < 1 {
		return 0, fmt.Errorf("invalid chunk size")
	}
	wg := sync.WaitGroup{}
	res := make(chan int, (len(nums)/сhunkSize)+1)
	var count int
	for i := 0; i+сhunkSize <= len(nums); i += сhunkSize {
		from := i
		to := from + сhunkSize
		wg.Add(1)
		count++
		go func() {
			defer wg.Done()
			summer(nums[from:to], res)
		}()
	}
	if len(nums)%сhunkSize != 0 {
		wg.Add(1)
		count++
		go func() {
			defer wg.Done()
			summer(nums[(len(nums)/сhunkSize)*сhunkSize:], res)
		}()
	}
	go func() {
		// ждем когда закончат все горутины и закроем канал
		wg.Wait()
		close(res)
	}()
	ans := 0
	for v := range res {
		ans += v
	}
	return ans, nil
}

func SumChunk(arr []int, result chan<- int) {
	var res int
	for _, elem := range arr {
		res += elem
	}
	result <- res
}
