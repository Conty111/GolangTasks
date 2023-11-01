/*
Map
Создайте дженерик функцию Map[T, U any](arr []T, transform func(T) U) []U, которая применяет
заданную функцию преобразования к каждому элементу среза и возвращает новый срез с преобразованными значениями.
Функция должна работать со срезами любого типа.

Примечания
Пример работы функции

arr := []int{1, 2, 3}
result := Map(arr, func(x int) string {
    return fmt.Sprintf("%d!", x)
})
fmt.Println(result) // Output: [1! 2! 3!]

*/

package generics

func Map[T, U any](arr []T, transform func(T) U) []U {
	res := make([]U, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = transform(arr[i])
	}
	return res
}
