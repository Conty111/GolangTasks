/*
Дан сервер доступный по адресу localhost:8082.
По запросу localhost:8082/mark?name=<имя студента> сервер возвращает: - код 200 и значение оценки студента,
если все прошло успешно - код 404, если студент не найден - код 500, если при с сервером случилась проблема

Напишите функцию Compare(name1, name2 string) (string, error),
которая сравнивает оценки двух студентов с именами name1 и name2 и выводит
> (оценка студента 1 больше оценки студента 2),
< (оценка студента 1 меньше оценки студента 2),
= (оценка студента 1 равна оценке студента 2).
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
	"sync"
)

type student struct {
	Note int
	Err  error
}

func Compare(name1, name2 string) (string, error) {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	var st1, st2 student
	go get_note(name1, wg, &st1)
	go get_note(name2, wg, &st2)
	wg.Wait()
	if st1.Err != nil {
		return "", st1.Err
	}
	if st2.Err != nil {
		return "", st2.Err
	}

	var res string
	if st1.Note > st2.Note {
		res = ">"
	} else if st1.Note < st2.Note {
		res = "<"
	} else {
		res = "="
	}
	return res, nil
}

func get_note(student_name string, wg *sync.WaitGroup, student *student) {
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
		student.Note = int(data[0])
	} else {
		if err != nil {
			student.Err = err
		} else {
			student.Err = fmt.Errorf("some error, request status is")
		}
	}
}
