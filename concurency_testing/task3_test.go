/*
Тест Reverse
Напишите тест для функции:

func ReverseString(input string) string {
    runes := []rune(input)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
Примечания
Функцию ReverseString реализовывать не нужно.
*/

package concurency_testing

import "testing"

func TestReverseString(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value string
		want  string
	}{
		{name: "default", value: "abc", want: "cba"},
		{name: "russian", value: "абвг", want: "гвба"},
		{name: "empty", value: "", want: ""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := ReverseString(tc.value)
			if got != tc.want {
				t.Errorf("ReverseString(%v); want %v, got %v", tc.value, tc.want, got)
			}
		})
	}
}
