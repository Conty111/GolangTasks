/*
Сортировка со слиянием
Даны два слайса. Напишите программу, содержащую функцию SortAndMerge(left, right []int) []int,
которая объединит слайсы в один отсортированный в два этапа:
- отсортировать каждый слайс
- объединить полученные слайсы в один.

Кстати, именно так работает алгоритм сортировки слиянием (merge sort)

Примечания
Ообъединять слайсы до сортировки не допустимо.
*/

package sortirovki

import (
	"slices"
)

func SortAndMerge(left, right []int) []int {
	res := make([]int, len(left)+len(right))
	slices.Sort(left)
	slices.Sort(right)
	var i, j, r int
	for {
		if j == len(right) {
			for idx := r; idx < len(res); idx++ {
				res[idx] = left[i]
				i++
			}
			break
		}
		if i == len(left) {
			for idx := r; idx < len(res); idx++ {
				res[idx] = right[j]
				j++
			}
			break
		}
		if left[i] < right[j] {
			res[r] = left[i]
			i++
		} else if left[i] > right[j] {
			res[r] = right[j]
			j++
		} else {
			res[r] = right[j]
			j++
			r++
			res[r] = left[i]
			i++
		}
		r++
	}
	return res
}
