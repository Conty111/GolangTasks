/*
Любимые студенты

Мы решили посчитать аналитику и посмотреть, а сколько же с нами учится каждый студент
курса по Go - то есть найти кол-во дней, которое он (студент) провел в курсе с момента
поступления и до 1 октября 2023 года.

Напишите функцию processJSON(jsonData []byte) error, которая должна принимать
данные о студентах в формате JSON, разбирать их и выводить искомое число дней.
Вывод должен быть в формате имяСтудента: количество дней

Формат ввода

[
    {"name": "Анна","admission_date": "2023-09-29T00:00:00Z"},
    {"name": "Иван","admission_date": "2023-09-28T00:00:00Z"}
]

Формат вывода

Анна: 2
Иван: 3

Примечания

type Student4 struct {
    Name         string    `json:"name"`
    AdmissionDate time.Time `json:"admission_date"`
    DaysOnCourse int       `json:"-"`
}
*/

package jsonstudy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type Student4 struct {
	Name          string    `json:"name"`
	AdmissionDate time.Time `json:"admission_date"`
	DaysOnCourse  int       `json:"-"`
}

func ProcessJSON(jsonData []byte) error {
	var s []Student4
	r := bytes.NewReader(jsonData)
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&s)
	if err != nil {
		return err
	}
	l, _ := time.LoadLocation("Europe/Vienna")
	date := time.Date(2023, time.Month(10), 1, 0, 0, 0, 0, l)
	for _, student := range s {
		fmt.Printf("%s: %d\n", student.Name, int(student.AdmissionDate.Sub(date).Abs().Hours()/24)+1)
	}
	return nil
}
