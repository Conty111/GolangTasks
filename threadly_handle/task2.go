/*
Напишите функцию NumbersGen(filename string) <-chan int,
которая будет читать файл и записывать в выходной канал числа из этого файла.
Формат файла: каждая строка содержит одно число, например:
2
4
6

Примечания
Если в строке не число, пропускайте значение.
*/

package threadly_handle

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

func NumbersGen1(filename string) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
		if err != nil {
			log.Println(err)
		}
		buf := bufio.NewReader(file)
		for {
			line, _, err := buf.ReadLine()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				} else {
					log.Println(err)
				}
			}
			num, err := strconv.Atoi(string(line))
			if err != nil {
				//log.Println(err)
				continue
			}
			out <- num
		}
	}()
	return out
}
