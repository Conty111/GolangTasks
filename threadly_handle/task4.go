/*
Реализуйте pipeline, который принимает на вход любое количество слайсов с числами,
выбирает из каждого положительные значения и возвращает произведение этих чисел.
Должны быть реализованы следующие функции:
- MultiplyPipeline(inputNums ...[]int) int
- NumbersGen(nums ...int) <-chan int
- Filter(in <-chan int) <-chan int
- Multiply(in <-chan int) int
*/

package threadly_handle

func MultiplyPipeline(inputNums ...[]int) int {
	res := 1
	for _, nums := range inputNums {
		res *= Multiply(Filter(NumbersGen(nums...)))
	}
	return res
}

func NumbersGen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, val := range nums {
			out <- val
		}
	}()
	return out
}

func Filter(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for val := range in {
			if val > 0 {
				out <- val
			}
		}
	}()
	return out
}

func Multiply(in <-chan int) int {
	res := 1
	for val := range in {
		res *= val
	}
	return res
}
