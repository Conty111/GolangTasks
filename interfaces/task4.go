/*
Напишите функцию Copy(r io.Reader, w io.Writer, n uint) error, которая копирует n байт из r в w.
Если количество байт, доступных для чтения меньше n - функция должна копировать все данные.
В случае ошибки - верните её.
*/
package interfaces

import "io"

func Copy(r io.Reader, w io.Writer, n uint) error {
	buf := make([]byte, n)
	readed, err := r.Read(buf)
	if err != nil {
		return err
	}
	_, err = w.Write(buf[:readed])
	if err != nil {
		return err
	}
	return nil
}
