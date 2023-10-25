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

import "fmt"

func MaxExpressionValue(nums []int) int {
	n := len(nums)
	arr_max := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		arr_max[i] = max(arr_max[i+1], nums[i])
	}
	fmt.Println(arr_max)
	arr_ans := make([]int, n+1)
	for i := n - 3; i >= 0; i-- {
		arr_ans[i] = max(arr_ans[i+1], arr_max[i+2]-nums[i+1]+arr_max[i+1]-nums[i])
	}
	return arr_ans[0]
}
