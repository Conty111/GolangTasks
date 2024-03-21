/*
Копирование

Дан слайс целых чисел nums. Этот слайс имеет емкость больше его длины.
Создайте функцию SliceCopy(nums []int) []int, которая вернёт новый слайс
длиной и ёмкостью, равной длине nums. Скопируйте в него значения из исходного слайса.
*/
package arraysandslices

func SliceCopy(nums []int) []int {
	length := len(nums)
	b := make([]int, length)
	for i := 0; i < length; i++ {
		b[i] = nums[i]
	}
	return b
}
