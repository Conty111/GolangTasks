/*
Дан сервер доступный по адресу localhost:8082.
По запросу localhost:8082/mark?name=<имя студента> сервер возвращает: - код 200 и значение оценки студента,
если все прошло успешно - код 404, если студент не найден - код 500, если при с сервером случилась проблема

Напишите функцию Average(names []string) (int, error),
которая выводит среднюю успеваемоть студентов с именами names.
Функция возвращает ошибку в случае невозможности получения оценки хотя бы одного студента.

Примечания
Используйте WaitGroup
*/

package sync_patterns

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type Student1 struct {
	Note int
	Err  error
}

func Average(names []string) (int, error) {
	students := make([]Student1, len(names))
	wg := &sync.WaitGroup{}
	wg.Add(len(names))
	for idx, elem := range names {
		go get_note1(elem, wg, &students[idx])
	}
	wg.Wait()
	var final_sum int
	for _, student := range students {
		if student.Err != nil {
			return final_sum, student.Err
		}
		final_sum += student.Note
	}
	return final_sum / len(names), nil
}

func get_note1(student_name string, wg *sync.WaitGroup, student *Student1) {
	defer wg.Done()
	cl := &http.Client{}
	url := fmt.Sprintf("http://localhost:8082/mark?name=%s", student_name)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		student.Err = err
		return
	}
	resp, err := cl.Do(req)
	if err != nil {
		student.Err = err
		return
	}
	defer resp.Body.Close()
	if (resp.StatusCode == 200 || resp.StatusCode == 201) && err == nil {
		data := make([]byte, resp.ContentLength)
		_, err := resp.Body.Read(data)
		if err != nil && !errors.Is(err, io.EOF) {
			student.Err = err
			return
		}
		note, err := strconv.Atoi(string(data))
		if err != nil {
			student.Err = err
			return
		}
		student.Note = note
	} else {
		if err != nil {
			student.Err = err
		} else {
			student.Err = fmt.Errorf("some error, request status is")
		}
	}
}
