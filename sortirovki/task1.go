/*

Сортировка слайса uint

Напишите программу, содержащую функцию SortNums(nums []uint), которая сортирует слайс nums по возрастанию

*/

package sortirovki

import "slices"

func SortNums(nums []uint) {
	slices.Sort(nums)
}
