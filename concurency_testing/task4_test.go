/*
Напишите тест для функции, которая проверяет, являются ли слова анограммами:

func AreAnagrams(str1, str2 string) bool {
    str1 = strings.ToLower(str1)
    str2 = strings.ToLower(str2)

    if len(str1) != len(str2) {
        return false
    }

    // Convert strings to slices of runes for sorting
    r1 := []rune(str1)
    r2 := []rune(str2)

    sort.Slice(r1, func(i, j int) bool {
        return r1[i] < r1[j]
    })

    sort.Slice(r2, func(i, j int) bool {
        return r2[i] < r2[j]
    })

    return string(r1) == string(r2)
}

*/

package concurency_testing

import "testing"

func TestAreAnagrams(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name   string
		value1 string
		value2 string
		want   bool
	}{
		{name: "positive", value1: "abc", value2: "cba", want: true},
		{name: "different length", value1: "abcc", value2: "abc", want: false},
		{name: "different letters", value1: "abcc", value2: "gasd", want: false},
		{name: "utf characters", value1: "авбг", value2: "бвга", want: true},
		{name: "empty", value1: "", value2: "", want: true},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := AreAnagrams(tc.value1, tc.value2)
			if got != tc.want {
				t.Errorf("AreAnagrams(%v, %v); want %v, got %v", tc.value1, tc.value2, tc.want, got)
			}
		})
	}
}
