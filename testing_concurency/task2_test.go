/*
Тест Contains numbers
Напишите тест для функции:

func Contains(numbers []int, target int) bool{
    for _, num := range numbers {
        if num == target {
            return true
        }
    }
    return false
}
Примечания
Функцию Contains реализовывать не нужно.

*/

package testing_concurency

import "testing"

func TestContains(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name   string
		arr    []int
		target int
		want   bool
	}{
		{name: "default", arr: []int{3, 2, 1, 4}, target: 1, want: true},
		{name: "default", arr: []int{3, 2, 1, 4}, target: 0, want: false},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := Contains(tc.arr, tc.target)
			if got != tc.want {
				t.Errorf("Contains(%v, %d); want %v, got %v", tc.arr, tc.target, tc.want, got)
			}
		})
	}
}
