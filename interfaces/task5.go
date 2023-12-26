/*
Напишите функцию Contains(r io.Reader, seq []byte) (bool, error)
которая должна найти первое вхождение байт seq в данных, доступных через Reader r.
Если последовательность найдена - верните true, nil, иначе false, nil.
В случае возникновения ошибки - false и возникшую ошибку.
*/
package interfaces

import (
	"errors"
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	current := make([]byte, len(seq)*2)
	g := true
	read, err := r.Read(current[:len(seq)])
	if err != nil {
		if errors.Is(err, io.EOF) {
			g = false
		} else {
			return false, err
		}
	}
	for g {
		read, err = r.Read(current[len(seq):])
		if err != nil {
			if errors.Is(err, io.EOF) {
				g = false
			} else {
				return false, err
			}
		}
		for i := 0; i < read; i++ {
			res := equal(current[i:len(seq)+i], seq)
			if res {
				return true, nil
			}
		}
		if read < len(seq) {
			break
		}
		tmp := current[len(seq):]
		for idx, val := range tmp {
			current[idx] = val
		}
	}
	return false, nil
}

func equal(fir, sec []byte) bool {
	isFinded := true
	for idx, val := range fir {
		if val != sec[idx] {
			isFinded = false
			break
		}
	}
	return isFinded
}
