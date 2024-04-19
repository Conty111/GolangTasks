/*
Значение выражения

Функция MaxExpressionValue(nums []int) int принимает на вход слайс nums.

Найдите максимальное значение выражения nums[s] — nums[r] + nums[q] — nums[p],
где p, q, r и s — индексы слайса, а s>r>q>p.

Например, для nums := []uint{3, 9, 10, 1, 30, 40} функция должна вернуть
значение 46 (поскольку 40 – 1 + 10 – 3 - максимально).

Задачу надо решить, используя принципы динамического программирования.
Примечания

В качестве решения надо отправить функцию MaxExpressionValue и все вспомогательные функции, которые вам потребуются.
*/
package dynamic

func MaxExpressionValue(nums []int) int {
	if len(nums) < 4 {
		return 0
	}
	first := make([]int, len(nums)+1)
	for i := len(nums) - 1; i >= 0; i-- {
		first[i] = max(first[i+1], nums[i]) // функция max - возвращает максимальное
	}
	second := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		second[i] = max(second[i+1], first[i+1]-nums[i])
	}
	arr := make([]int, len(nums)+1)
	var tmp int = nums[0]
	for i := 1; i < len(nums)+1; i++ {
		arr[i] = max(arr[i-1], nums[i-1]-tmp)
		if tmp > nums[i-1] {
			tmp = nums[i-1]
		}
	}
	arr = arr[1:]
	var res int
	for i := 1; i < len(nums)-2; i++ {
		res = max(res, arr[i]+second[i+1])
	}
	return res
}
