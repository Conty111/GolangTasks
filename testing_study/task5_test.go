/*
Фукнция GetUTFLength(input []byte) (int, error) возвращает длину строки UTF8 и ошибку ErrInvalidUTF8 (в случае возникновения).
Напишите тест, который бы проверял возвращаемые функцией значения.

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
if !utf8.Valid(input) {
return 0, ErrInvalidUTF8
}

return utf8.RuneCount(input), nil
}

Примечания
Функцию GetUTFLength реализовывать не нужно.
*/

package testing_study

import (
	"errors"
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	cases := []struct {
		name      string
		value     []byte
		want      int
		wantError error
	}{
		{name: "default", value: []byte("aboba"), want: 5},
		{name: "russian symbols", value: []byte("август"), want: 6},
		{name: "with invalid bytes", value: []byte{0xff, 0xfe, 0xfd}, want: 0, wantError: ErrInvalidUTF8},
		{name: "with not existing symbols", value: []byte{254, 255, 255}, want: 0, wantError: ErrInvalidUTF8},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetUTFLength(tc.value)
			if err != nil {
				if !errors.Is(err, tc.wantError) {
					t.Errorf("GetUTFLength(%v) returns not existing error: got: %v, existing: %v", tc.value, err, tc.wantError)
				}
			} else if got != tc.want {
				t.Errorf("GetUTFLength(%v) = %v; want %v", tc.value, got, tc.want)
			}
		})
	}
}
