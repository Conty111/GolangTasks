/*
Функция SortIntegers(numbers []int) (пакет main) сортирует переданный слайс в порядке увеличения значений.
Напишите тест для проверки корректности работы.

Примечания
Функцию SortIntegers реализовывать не нужно.
*/

package concurency_testing

import (
	"slices"
	"testing"
)

func TestSortIntegers(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value []int
		want  []int
	}{
		{name: "default", value: []int{3, 2, 1, 4}, want: []int{1, 2, 3, 4}},
		{name: "negative nums", value: []int{-3, 2, -1, 4}, want: []int{-3, -1, 2, 4}},
		{name: "empty", value: []int{}, want: []int{}},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			SortIntegers(tc.value)
			if !slices.Equal(tc.value, tc.want) {
				t.Errorf("SortIntegers(%v) = %v; want %v", tc.value, tc.value, tc.want)
			}
		})
	}
}
