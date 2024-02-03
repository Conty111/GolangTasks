/*
Напишите программу для чтения JSON-файла и отображения содержимого в терминале,
используйте концепцию таймаута вместе с контекстом и отменой при чтении JSON-файла.
Реализуйте функцию
func readJSON(ctx context.Context, path string, result chan<- []byte)
ctx - контекст
path - путь к json
result - канал, в который нужно вывести прочитанное значение
*/

package timeouts_pattern

import (
	"bufio"
	"context"
	"errors"
	"io"
	"log"
	"os"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	defer close(result)
	file, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	buf := make([]byte, 10)
	reader := bufio.NewReader(file)
	resArr := make([]byte, 0)
	for {
		select {
		case <-ctx.Done():
			break
		default:
			n, err := reader.Read(buf)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					log.Println(err)
				} else {
					resArr = append(resArr, buf[:n]...)
				}
				break
			}
			resArr = append(resArr, buf[:n]...)
		}
	}
	result <- resArr
}
