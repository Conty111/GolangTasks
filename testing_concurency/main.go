package testing_concurency

import (
	"slices"
	"sort"
	"strings"
)

func SortIntegers(numbers []int) {
	slices.Sort(numbers)
}

func Contains(numbers []int, target int) bool {
	for _, num := range numbers {
		if num == target {
			return true
		}
	}
	return false
}

func ReverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

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
