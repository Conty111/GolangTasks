/*
Слияние двух частей

Дан слайс nums, состоящий из 2n элементов в формате [x0,x1,...,xn,y0,y1,...,yn].
Создайте функцию Mix(nums []int) []int, которая вернёт слайс,
содержащий значения в следующем порядке: [x0,y0,x1,y1,...,xn,yn].
*/
package arraysandslices

func Mix(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	for i := 0; i < l; i += 2 {
		res[i] = nums[i/2]
		res[i+1] = nums[i/2+l/2]
	}
	return res
}
