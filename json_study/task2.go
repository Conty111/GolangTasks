/*
Колобок

По амбарам помела, по сусекам поскребла — базу данных получила.
Напишите функцию mergeJSONData(jsonDataList ...[]byte) ([]byte, error),
которая принимает несколько экземпляров данных в формате JSON, объединяет их в один экземпляр и возвращает его.
Примечания

Например: В функцию передаются два JSON:

[
    {"name": "Oleg","class": "9B"},
    {"name": "Ivan","class": "9A"}
]
и
[
    {"name": "Maria","class": "9B"},
    {"name": "John","class": "9A"}
]

На выходе нужно получить:

[
    {"class": "9B","name": "Oleg"},
    {"class": "9A","name": "Ivan"},
    {"class": "9B","name": "Maria"},
    {"class": "9A","name": "John"}
]
*/

package jsonstudy

import (
	"bytes"
	"encoding/json"
)

type Student2 struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func SplitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	var s []Student2
	r := bytes.NewReader(jsonData)
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&s)
	if err != nil {
		return nil, err
	}
	info := make(map[string][]byte)
	tmp := make(map[string][]Student2)
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
