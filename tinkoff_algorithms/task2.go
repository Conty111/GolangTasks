/*
Поворот матрицы на 90 градусов

Примеры данных
Пример 1
2 2
1 1
2 3

2 1
3 1

Пример 2
2 3
1 2 3
4 5 6
2 3
1 2 3
4 5 6
4 1
5 2
6 3
4 1
5 2
6 3
Пример 3
1 1
69
1 1
69
69

*/

package tinkoff_algorithms

import (
	"fmt"
)

func Task2() {
	var rows, cols int
	fmt.Scan(&rows, &cols)
	//mx := make([][]int, n)
	mx := make([][]int, rows)
	for i := 0; i < rows; i++ {
		mx[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			fmt.Scan(&mx[i][j])
		}
	}

	// Создаем новую матрицу для хранения повернутой матрицы
	newMatrix := make([][]int, cols)
	for i := 0; i < cols; i++ {
		newMatrix[i] = make([]int, rows)
	}

	// Поворачиваем исходную матрицу на 90 градусов
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			newMatrix[j][rows-1-i] = mx[i][j]
		}
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < rows-1; j++ {
			fmt.Printf("%d ", newMatrix[i][j])
		}
		fmt.Println(newMatrix[i][rows-1])
	}
}
