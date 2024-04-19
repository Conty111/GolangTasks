/*
Напишите функцию ReadCSV(file string) (<-chan []string, error) для чтения csv-файлов.
В выходной канал должны передаваться строки из файла. Если возникла ошибка, верните ее описание.
*/

package threadly_handle2

import (
	"encoding/csv"
	"os"
)

func ReadCSV(file string) (<-chan []string, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	buf := csv.NewReader(f)
	out := make(chan []string)
	go func() {
		defer close(out)
		for {
			line, err := buf.Read()
			if err != nil {
				if line != nil {
					out <- line
				}
				return
			}
			out <- line
		}
	}()
	return out, nil
}
