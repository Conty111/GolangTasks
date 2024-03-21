/*
	Суммирование

Напишите универсальную функцию Sum[T Number](arr []T) T, которая может суммировать элементы
срезов разных числовых типов (например, int, float64). Функция должна брать фрагмент любого
числового типа и возвращать его сумму.

Примечания
Напишите так же констрейнт type Number interface, который будет обозначать все численные типы.
*/
package generics

type Number interface {
	int | int8 | int16 | int32 | int64 | float64 | float32 | uint | uint8 | uint16 | uint32 | uint64
}

func Sum[T Number](arr []T) T {
	var res T
	for _, elem := range arr {
		res += elem
	}
	return res
}
