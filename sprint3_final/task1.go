package sprint3_final

import (
	"encoding/json"
	"reflect"
)

func CompareJSON(json1, json2 []byte) (bool, error) {
	var obj1, obj2 interface{}

	err := json.Unmarshal(json1, &obj1)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(json2, &obj2)
	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(obj1, obj2), nil
}
