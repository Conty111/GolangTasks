package testing_study

import (
	"errors"
	"slices"
	"unicode/utf8"
)

func Sum(a, b int) int {
	return a + b
}

func Length(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "short"
	case a < 100:
		return "long"
	}
	return "very long"
}

func Multiply(a, b int) int {
	return a * b
}

func DeleteVowels(s string) string {
	vowels := []rune{'a', 'o', 'e', 'u', 'i'}
	newString := make([]rune, 0)
	for _, letter := range s {
		if !slices.Contains(vowels, letter) {
			newString = append(newString, letter)
		}
	}
	return string(newString)
}

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}
