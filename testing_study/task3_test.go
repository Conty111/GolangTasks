/*
Тест Multipy
Функция Multiply(a, b int) int (пакет main) возвращает произведение двух переданных чисел.
Напишите тест TestMultiply(t *testing.T) для проверки корректности работы.

Примечания
Функцию Multipy(a, b int) int реализовывать не нужно.
*/

package testing_study

import "testing"

func TestMultiply(t *testing.T) {
	cases := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive", a: 1, b: 2, want: 2},
		{name: "zero", a: 0, b: 2, want: 0},
		{name: "negative nums", a: -2, b: 2, want: -4},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := Multiply(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Multiply(%d, %d) = %v; want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
