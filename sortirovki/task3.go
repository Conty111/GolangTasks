/*
Сортировка слайса строк

Напишите программу, содержащую функцию `SortNames(names []string)``,
которая сортирует список имён в алфавитном порядке.
Если первые символы совпадают, сортировать по вторым, и т.д.
Примечания

Пример отсортированного списка: Аксинья, Арина, Варвара, Есения
*/

package sortirovki

import "slices"

func SortNames(names []string) {
	slices.SortFunc(names, func(a, b string) int {
		for i := 0; i < min(len(a), len(b)); i++ {
			if a[i] < b[i] {
				return -1
			} else if a[i] > b[i] {
				return 1
			}
		}
		if len(a) > len(b) {
			return 1
		} else if len(a) < len(b) {
			return -1
		}
		return 0
	})
}
