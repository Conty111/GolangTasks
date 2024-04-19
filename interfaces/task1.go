/*
Напишите функцию WriteString(s string, w io.Writer) error,
которая записывает строку s в место назначения, используя интерфейс w.
В случае возникновения ошибки - функция должна её возвращать, иначе nil.
*/

package interfaces

import "io"

func WriteString(s string, w io.Writer) error {
	_, err := w.Write([]byte(s))
	return err
}
