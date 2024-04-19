/*
Дан сервер доступный по адресу localhost:8082. По запросу localhost:8082/mark?name=<имя студента>
сервер возвращает: - код 200 и значение оценки студента, если все прошло успешно - код 404,
если студент не найден - код 500, если при с сервером случилась проблема

Напишите функцию CompareList(names []string) (map[string]string, error),
выводит карту, где ключом является имя студента из списка name,
а значением является > (оценка студента больше средней оценки студентов),
< (оценка студента меньше средней оценки студентов),
= (оценка студента равна средней оценки студентов).
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

type Student3 struct {
	Name string
	Note float32
	Err  error
}

func CompareList(names []string) (map[string]string, error) {
	students := make([]Student3, len(names))
	wg := &sync.WaitGroup{}
	wg.Add(len(names))
	for idx, elem := range names {
		go get_note3(elem, wg, &students[idx])
	}
	wg.Wait()
	var avg_sum float32
	for _, s := range students {
		if s.Err != nil {
			return nil, s.Err
		}
		avg_sum += s.Note
	}
	avg := avg_sum / float32(len(students))
	res := make(map[string]string)
	for i := 0; i < len(students); i++ {
		var znak string
		if students[i].Note > avg {
			znak = ">"
		} else if students[i].Note < avg {
			znak = "<"
		} else {
			znak = "="
		}
		res[students[i].Name] = znak
	}
	return res, nil
}

func get_note3(student_name string, wg *sync.WaitGroup, student *Student3) {
	defer wg.Done()
	student.Name = student_name
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
		student.Note = float32(data[0])
	} else {
		if err != nil {
			student.Err = err
		} else {
			student.Err = fmt.Errorf("some error, request status is")
		}
	}
}
