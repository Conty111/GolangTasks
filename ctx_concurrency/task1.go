/*
Напишите функцию Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error)
которая должна найти первое вхождение байт seq в данных, доступных через Reader r.
Если последовательность найдена - верните true, nil, иначе false, nil. В случае возникновения ошибки - false и возникшую ошибку.
В случае отмены контекста - функция должна вернуть false и ошибку - причину отмены контекста.
*/

package ctx_concurrency

import (
	"context"
	"errors"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	tmp := make([]byte, 1)
	pos := 0
	n := len(seq)
	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
			_, err := r.Read(tmp)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					return false, err
				}
				return false, nil
			}
			if tmp[0] == seq[pos] {
				pos += 1
			} else {
				pos = 0
			}
			if pos == n-1 {
				return true, nil
			}
		}
	}
}
