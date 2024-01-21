/*
Реализуйте pipeline, который читает файл с числами (каждая строка - число),
игнорирует нечетные числа, а чётные суммирует.
Для этого напишите функцию
- SumValuesPipeline(filename string) int, вызывающую:
- NumbersGen(filename string) <-chan int для чтения файла
- Filter(in <-chan int) <-chan int - для фильтрации чисел
- Sum(in <-chan int) int - для суммирования результата.
Напишите реализацию перечисленных функций
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

func SumValuesPipeline(filename string) int {
	return Sum(Filter2(NumbersGen2(filename)))
}

func NumbersGen2(filename string) <-chan int {
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
				continue
			}
			out <- num
		}
	}()
	return out
}

func Filter2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for val := range in {
			if val%2 == 0 {
				out <- val
			}
		}
	}()
	return out
}

func Sum(in <-chan int) int {
	var res int
	for val := range in {
		res += val
	}
	return res
}
