/*
Сортировка символов по частоте
Дана строка с символами из набора алфавита. Напишите программу с функцией SortByFreq(str string) string,
которая будет сортировать символы из строки по возрастанию с учетом частоты повторения.
Символы с наименьшим количеством вхождений должны идти в начале, а символы с наибольшей частотой - в конце.
В случае одинаковой частоты символов, они должны быть отсортированы в алфавитном порядке.

Примечания
Пример:
Вход: "abbbzzzat"
Выход: "taabbbzzz"
*/

package sortirovki

import (
	"slices"
)

func SortByFreq(str string) string {
	res := []byte(str)
	letters := make(map[byte]int, 0)
	for _, elem := range res {
		letters[elem] += 1
	}
	slices.SortFunc(res, func(a, b byte) int {
		if letters[a] < letters[b] {
			return -1
		} else if letters[a] > letters[b] {
			return 1
		} else {
			if a > b {
				return 1
			} else if a < b {
				return -1
			}
			return 0
		}
	})
	return string(res)
}
