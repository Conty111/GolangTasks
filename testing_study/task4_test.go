/*
Напишите тест для функции DeleteVowels(s string) string,
которая должна удалять все гласные из строки английского языка (y не считается гласной).
Используйте table driven testing.

Примечания
Функцию DeleteVowels реализовывать не нужно.
*/

package testing_study

import "testing"

func TestDeleteVowels(t *testing.T) {
	cases := []struct {
		name  string
		value string
		want  string
	}{
		{name: "default", value: "aboba", want: "bb"},
		{name: "with y", value: "yes", want: "ys"},
		{name: "without vowels", value: "nnnn", want: "nnnn"},
		{name: "empty", value: "", want: ""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := DeleteVowels(tc.value)
			if got != tc.want {
				t.Errorf("DeleteVowels(%s) = %v; want %v", tc.value, got, tc.want)
			}
		})
	}
}
