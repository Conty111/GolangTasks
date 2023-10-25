/*
Лимит по значению

Дан неотсортированный слайс целых чисел. Напишите функцию UnderLimit(nums []int, limit int, n int) ([]int, error),
которая будет возвращать первые n (либо меньше, если остальные не подходят)
элементов, которые меньше limit. В случае ошибки функция должна вернуть nil и описание ошибки.
*/
package arraysandslices

import "fmt"

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if n <= 0 || nums == nil || limit == 0 {
		return nil, fmt.Errorf("wrong data")
	}
	res := make([]int, 0, n)
	for _, elem := range nums {
		if elem < limit {
			res = append(res, elem)
			if len(res) == cap(res) {
				break
			}
		}
	}
	return res, nil
}
