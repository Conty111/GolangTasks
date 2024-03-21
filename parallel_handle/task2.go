/*
Параллельная обработка WordCounter
Напишите программу, которая читает файлы параллельно и считывает количество вхождений слов во всех файлах.
Код программы должен содержать объявление следующей структуры:

type WordCounter struct {
	wordsCount map[string]int // здесь должен быть список слов с указанием количества повторений во всех файлах.
	// можно добавлять другие поля
}

WordCounter должен удовлетворять следующему интерфейсу:

type CounterWorker interface{
	ProcessFiles(files ...string) error // для запуска обработки файлов
	ProcessReader(r io.Reader) error // для подсчёта слов в одном файле
}

Примечания
Cчитайте, что текст не содержит знаков пунктуации, то есть за слово принимайте любую единицу текста, обрамлённую пробелами.
*/

package gzipper

import (
	"io"
	"os"
	"strings"
	"sync"
)

type CounterWorker interface {
	ProcessFiles(files ...string) error // для запуска обработки файлов
	ProcessReader(r io.Reader) error    // для подсчёта слов в одном файле
}

type WordCounter struct {
	wordsCount map[string]int // здесь должен быть список слов с указанием количества повторений во всех файлах.
	// можно добавлять другие поля
	mu sync.Mutex
}

func (w *WordCounter) ProcessFiles(files ...string) error {
	wg := sync.WaitGroup{}
	w.mu = sync.Mutex{}
	w.wordsCount = make(map[string]int)
	for _, fileName := range files {
		file, err := os.OpenFile(fileName, os.O_RDONLY, 0777)
		if err != nil {
			return err
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := w.ProcessReader(file); err != nil {
				return
			}
		}()
	}
	wg.Wait()
	return nil
}

func (w *WordCounter) ProcessReader(r io.Reader) error {
	all, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	for _, word := range strings.SplitN(string(all), " ", -1) {
		word = strings.ToLower(word)
		if len(word) > 0 {
			w.mu.Lock()
			_, ok := w.wordsCount[word]
			if !ok {
				w.wordsCount[word] = 1
			} else {
				w.wordsCount[word] += 1
			}
			w.mu.Unlock()
		}
	}
	return nil
}

func (w *WordCounter) GetWords() map[string]int {
	return w.wordsCount
}
