/*
С новым учебным годом!

1 сентября каждого учебного года во всех базах данных школьников происходит великий пересчёт.
Напишите функцию modifyJSON(jsonData []byte) ([]byte, error), которая принимает данные в формате JSON,
добавляет 1 год к полю grade и возвращает обновлённые данные также в формате JSON

Формат ввода
[{"name": "Oleg","grade": 9},{"name": "Katya","grade": 10}]

Формат вывода
[{"name": "Oleg","grade": 10},{"name": "Katya","grade": 11}]

Примечания:
Структура ученика

type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}
*/

package jsonstudy

import (
	"encoding/json"
	"strings"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func ModifyJSON(jsonData []byte) ([]byte, error) {
	if string(jsonData) == "null" {
		return nil, nil
	}
	var s []Student
	buf := strings.NewReader(string(jsonData))
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&s)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(s); i++ {
		s[i].Grade += 1
	}
	res, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return res, nil
}
