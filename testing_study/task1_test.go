/*
Тест Sum
Функция Sum(a, b int) int (пакет main) возвращает результат суммирования чисел a и b.
Напишите тест TestSum(t *testing_study.T) для проверки корректности работы.

Примечания
Функцию Sum(a, b int) int реализовывать не нужно.
*/

package testing_study

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "positive values",
			a:    1,
			b:    1,
			want: 2,
		},
		{
			name: "mixed values",
			a:    -1,
			b:    1,
			want: 0,
		},
		{
			name: "long values",
			a:    100000000000000000,
			b:    1000000000000000000,
			want: 1100000000000000000,
		},
	}
	for _, tc := range cases {
		tc := tc
		// запуск отдельного теста
		t.Run(tc.name, func(t *testing.T) {
			// тестируем функцию Sum
			got := Sum(tc.a, tc.b)
			// проверим полученное значение
			if got != tc.want {
				t.Errorf("Sum(%d, %d) = %v; want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
