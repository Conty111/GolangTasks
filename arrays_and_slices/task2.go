/*
Удаление элемента

Дан неотсортированный слайс целых чисел.
Напишите функцию Clean(nums []int, x int) ([]int), которая удаляет из исходного слайса все вхождения x.
Важно - не использовать дополнительный слайс.
*/
package arraysandslices

func Clean(nums []int, x int) []int {
	var cleaned int
	length := len(nums)
	for i := 0; i < length-cleaned; i++ {
		if nums[i] == x {
			for j := i; j < length-1; j++ {
				nums[j] = nums[j+1]
			}
			cleaned++
		}
	}
	if cleaned == length-1 && nums[length-cleaned-1] == x {
		return nums[:length-cleaned-1]
	}
	return nums[:length-cleaned]
}
