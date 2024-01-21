/*
Параллельная обработка gzipper
Напишите пакет (gzipper) для сжатия файлов с помощью compress/gzip. Для этого реализуйте следующие функции:
- FileNameGen(dir string, pattern *regexp.Regexp) <-chan Work для получения файлов в директории по заданному регулярному выражению
- compress(jobs <-chan Work) для сжатия файлов из канала и записи результатов на диск.
Имена сжатых файлов должны быть сформированы по правилу имяисходногофайла.gz,
например для файла myfile.txt -> myfile.txt.gz. Используйте горутины для параллельной обработки файлов.

Примечания
Код программы должен содержать объявление следующей структуры:
type Work struct {
	FilePath string
}
*/

package gzipper

import (
	"compress/gzip"
	"io"
	"os"
	"regexp"
)

type Work struct {
	FilePath string
}

func FileNameGen(dir string, pattern *regexp.Regexp) <-chan Work {
	out := make(chan Work)
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	go func() {
		defer close(out)
		for _, file := range files {
			if !file.IsDir() {
				if len(pattern.Find([]byte(file.Name()))) > 0 {
					out <- Work{FilePath: dir + "/" + file.Name()}
				}
			}
		}
	}()
	return out
}

func compress(jobs <-chan Work) {
	for w := range jobs {
		file, err := os.OpenFile(w.FilePath, os.O_RDWR, 0777)
		defer file.Close()
		if err != nil {
			return
		}
		newFile, err := os.Create(w.FilePath + ".gz")
		if err != nil {
			panic(err)
		}
		defer newFile.Close()

		// Create a new gzip writer
		gzipWriter := gzip.NewWriter(newFile)
		defer gzipWriter.Close()

		// Copy the contents of the original file to the gzip writer
		_, err = io.Copy(gzipWriter, file)
		if err != nil {
			panic(err)
		}

		// Flush the gzip writer to ensure all data is written
		gzipWriter.Flush()
	}
}
