/*
Дженерик стэк
Реализуйте дженерик структуру данных стека, которая может извлекать и извлекать элементы любого типа.
Функции, которые нужно реализовать Push(val T) Pop() T

Примечания
Структура должна иметь вид Stack[T any] и иметь конструктор NewStack.
*/

package generics

type Stack[T any] struct {
	length int
	arr    []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		length: 0,
		arr:    make([]T, 0),
	}
}

func (s *Stack[T]) Push(val T) {
	s.arr = append(s.arr, val)
	s.length += 1
}

func (s *Stack[T]) Pop() T {
	var res T
	res = s.arr[s.length-1]
	s.length -= 1
	s.arr = s.arr[:s.length]
	return res
}
