/*

Функция Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error)
находит первое вхождение байт seq в данных, доступных через Reader r.
Если последовательность найдена - она возвращает true, nil, иначе false, nil.
В случае возникновения ошибки - false и возникшую ошибку.
В случае отмены контекста - функция возвращает false и ошибку - причину отмены контекста.

Напишите тесты для проверки корректности работы.
Примечания

Покрытие кода функции должно быть 100%.

*/

package sprint3_final

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"
	"time"
)

type CustomReader struct{}

func (cr *CustomReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("custom error")
}

func TestContains(t *testing.T) {
	defaultCases := []struct {
		name      string
		r         io.Reader
		seq       []byte
		want      bool
		wantError bool
	}{
		{
			name:      "default positive test",
			r:         bytes.NewReader([]byte{0, 1, 2, 3, 4}),
			seq:       []byte{1, 2, 3},
			want:      true,
			wantError: false,
		},
		{
			name:      "sub sequence equals data",
			r:         bytes.NewReader([]byte{0, 1, 2, 3, 4}),
			seq:       []byte{0, 1, 2, 3, 4},
			want:      true,
			wantError: false,
		},
		{
			name:      "empty sequence",
			r:         bytes.NewReader([]byte{0, 1, 2, 3, 4}),
			seq:       make([]byte, 0),
			want:      true,
			wantError: false,
		},
		{
			name:      "empty source and seq",
			r:         bytes.NewReader(make([]byte, 0)),
			seq:       make([]byte, 0),
			want:      false,
			wantError: false,
		},
		{
			name:      "default negative",
			r:         bytes.NewReader([]byte{0, 1, 2, 3, 4}),
			seq:       []byte{0, 1, 2, 2},
			want:      false,
			wantError: false,
		},
		{
			name:      "sequence is bigger than source",
			r:         bytes.NewReader([]byte{0, 1, 2, 3, 4}),
			seq:       []byte{0, 1, 2, 3, 4, 5},
			want:      false,
			wantError: false,
		},
		{
			name:      "with existing error",
			r:         bytes.NewReader([]byte{0, 1, 2, 3, 4}),
			seq:       []byte{0, 1, 2, 3, 4, 5},
			want:      false,
			wantError: false,
		},
	}

	for _, tc := range defaultCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()
			got, err := Contains(ctx, tc.r, tc.seq)
			if err != nil && !tc.wantError {
				t.Errorf("Contains(%v, %v, %v) returns not existing error: got: %v", ctx, tc.r, tc.seq, err)
			} else if got != tc.want {
				t.Errorf("Contains(%v, %v, %v) = %v; want %v", ctx, tc.r, tc.seq, got, tc.want)
			}
		})
	}

	t.Run("with context timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Nanosecond*5)
		defer cancel()
		size := 1024
		data := make([]byte, size)
		for i := 0; i < size; i++ {
			data[i] = byte(i)
		}
		_, err := Contains(ctx, bytes.NewReader(data), data[15:32])
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("with big data", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.TODO())
		defer cancel()
		size := 102400
		data := make([]byte, size)
		for i := 0; i < size; i++ {
			data[i] = byte(i)
		}
		got, err := Contains(ctx, bytes.NewReader(data), data[150:322])
		if err != nil {
			t.Errorf("Contains() returns not existing error: got: %v", err)
		} else if got != true {
			t.Errorf("Contains() returns %v; want %v", got, true)
		}
	})

	t.Run("invalid reader", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.TODO())
		defer cancel()
		var reader CustomReader
		_, err := Contains(ctx, &reader, []byte{1})
		if err == nil {
			t.Errorf("expected error: got nil")
		}
	})

}
