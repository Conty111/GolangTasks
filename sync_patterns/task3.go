/*
Дан сервер доступный по адресу localhost:8082.
По запросу localhost:8082/mark?name=<имя студента> сервер возвращает: - код 200 и значение оценки студента,
если все прошло успешно - код 404, если студент не найден - код 500, если при с сервером случилась проблема

Напишите функцию BestStudents(names []string) (string, error),
выводяющу список студентов с оценками выше средней успеваемости студентов из списка names в алфавитном порядке через запятую.
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
	"slices"
	"sync"
)

type Student struct {
	Name string
	Note float32
	Err  error
}

func BestStudents(names []string) (string, error) {
	students := make([]Student, len(names))
	wg := &sync.WaitGroup{}
	wg.Add(len(names))
	for idx, elem := range names {
		go get_note2(elem, wg, &students[idx])
	}
	wg.Wait()
	var avg_sum float32
	for _, s := range students {
		if s.Err != nil {
			return "", s.Err
		}
		avg_sum += s.Note
	}
	avg := avg_sum / float32(len(students))
	var res []string
	for i := 0; i < len(students); i++ {
		if students[i].Note > avg {
			res = append(res, students[i].Name)
		}
	}
	slices.Sort(res)
	str_res := res[0]
	if len(res) > 1 {
		for _, elem := range res[1:] {
			str_res += fmt.Sprintf(", %s", elem)
		}
	}
	return str_res, nil
}

func get_note2(student_name string, wg *sync.WaitGroup, student *Student) {
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
