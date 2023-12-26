/*
Reader_Writer_2
Напишите функцию ReadString(r io.Reader) (string, error),
которая читает данные с помощью r и возвращает эти данные в строковом виде.
В случае возникновения ошибки функция должна вернуть пустую строку и возникшую ошибку,
иначе строку и nil.
*/
package interfaces

import (
	"errors"
	"io"
)

func ReadString(r io.Reader) (string, error) {
	bufSize := 10
	var res []byte
	p := make([]byte, bufSize)
	n, err := r.Read(p)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return "", nil
		}
		return "", err
	}
	res = append(res, p[:n]...)
	for n == bufSize {
		n, err = r.Read(p)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return "", err
		}
		res = append(res, p[:n]...)
	}
	return string(res), nil
}
