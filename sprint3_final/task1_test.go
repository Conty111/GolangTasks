/*


Функция CompareJSON(json1, json2 []byte) (bool, error), принимает на вход два
объекта json и сравнивает их.
Если они равны, то функция возвращает true, nil, иначе false, nil, либо описание ошибки.
Напишите тесты для проверки корректности работы.
Примечания

- порядок следования полей в json в равных объектах может быть разный
- json не содержит вложенных объектов
- покрытие кода функции должно быть 100%.

*/

package sprint3_final

import (
	"testing"
)

func TestCompareJSON(t *testing.T) {
	cases := []struct {
		name      string
		value1    []byte
		value2    []byte
		want      bool
		wantError bool
	}{
		{name: "default positive test",
			value1: []byte(`{"key": "value"}`), value2: []byte(`{"key": "value"}`), want: true, wantError: false},
		{name: "pairs in different order",
			value1: []byte(`{"key1": "1", "key2": 2}`), value2: []byte(`{"key2": 2, "key1": "1"}`), want: true, wantError: false},
		{name: "with enclosure",
			value1: []byte(`{"key": {"en_key": "ha"}}`), value2: []byte(`{"key": {"en_key": "ha"}}`), want: true, wantError: false},

		{name: "different values",
			value1: []byte(`{"key": "value"}`), value2: []byte(`{"key": "value1"}`), want: false, wantError: false},
		{name: "different keys",
			value1: []byte(`{"key1": "value"}`), value2: []byte(`{"key": "value"}`), want: false, wantError: false},
		{name: "different value types",
			value1: []byte(`{"key": "1"}`), value2: []byte(`{"key": 1}`), want: false, wantError: false},
		{name: "different value types",
			value1: []byte(`{"key": ["1"]}`), value2: []byte(`{"key": "1"}`), want: false, wantError: false},

		{name: "invalid bytes",
			value1: []byte(`{"key": ["1"]}a`), value2: []byte(`{"key": "1"}a`), want: false, wantError: true},
		{name: "invalid json",
			value1: []byte(`{"key": ["1"]}`), value2: []byte(`{key: "1"}a`), want: false, wantError: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := CompareJSON(tc.value1, tc.value2)
			if err != nil && !tc.wantError {
				t.Errorf("CompareJSON(%v, %v) returns not existing error: got: %v, existing: %v", tc.value1, tc.value2, err, tc.wantError)
			} else if got != tc.want {
				t.Errorf("CompareJSON(%v, %v) = %v; want %v", tc.value1, tc.value2, got, tc.want)
			}
		})
	}

}
