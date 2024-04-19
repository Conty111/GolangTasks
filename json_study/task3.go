/*
Равнение на флаг!

На линейке ученикам нужно сгруппироваться по классам.
Проведём линейку для базы данных.

Напишите функцию splitJSONByClass3(jsonData []byte) (map[string][]byte, error),
которая принимает данные в формате JSON и возвращает мапу, в которой ключи — классы,
а значения — данные в формате JSON, которые к этому классу относятся.

Например: Входные данные
[
    {"name": "Oleg","class": "9B"},
    {"name": "Ivan","class": "9A"},
    {"name": "Maria","class": "9B"},
    {"name": "John","class": "9A"}
]

Выходные данные должны быть в виде map:

map[string][]byte{
        "9A": []byte(`[{"class":"9A","name":"Ivan"},{"class":"9A","name":"John"}]`),
        "9B": []byte(`[{"class":"9B","name":"Oleg"},{"class":"9B","name":"Maria"}]
}
*/

package jsonstudy

import (
	"bytes"
	"encoding/json"
)

type Student3 struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func SplitJSONByClass3(jsonData []byte) (map[string][]byte, error) {
	var s []Student3
	r := bytes.NewReader(jsonData)
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&s)
	if err != nil {
		return nil, err
	}
	info := make(map[string][]byte)
	tmp := make(map[string][]Student3)
	for _, elem := range s {
		tmp[elem.Class] = append(tmp[elem.Class], elem)
	}
	for key, val := range tmp {
		info[key], err = json.Marshal(val)
		if err != nil {
			return nil, err
		}
	}
	return info, nil
}
