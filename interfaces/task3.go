/*
Создайте структуру UpperWriter, содержащую поле UpperString string.
Реализуйте интерфейс io.Writer. Метод Write должен переводить строку
в верхний регистр и записывать данные в поле UpperString.
В случае ошибки - верните её.
*/
package interfaces

import "strings"

type UpperWriter struct {
	UpperString string
}

func (uw *UpperWriter) Write(p []byte) (n int, err error) {
	uw.UpperString = strings.ToUpper(string(p))
	return len(p), nil
}
