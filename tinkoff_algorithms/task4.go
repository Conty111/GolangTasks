package tinkoff_algorithms

import (
	"fmt"
	"strings"
)

func rotateRight(matrix [][]int) []string {
	steps := make([]string, 0)
	// Транспонировать матрицу
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			steps = append(steps, fmt.Sprintf("%d %d %d %d\n", i, j, j, i))
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Перевернуть каждую строку матрицы
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			steps = append(steps, fmt.Sprintf("%d %d %d %d\n", i, j, i, len(matrix)-1-j))
			matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]

		}
	}
	return steps
}

func rotateLeft(matrix [][]int) []string {
	steps := make([]string, 0)
	// Транспонировать матрицу
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			steps = append(steps, fmt.Sprintf("%d %d %d %d\n", i, j, j, i))
		}
	}

	// Перевернуть каждый столбец матрицы
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[i][j], matrix[len(matrix)-1-i][j] = matrix[len(matrix)-1-i][j], matrix[i][j]
			steps = append(steps, fmt.Sprintf("%d %d %d %d\n", i, j, len(matrix)-1-i, j))
		}
	}
	return steps
}

func Task4() {
	var n int
	var direction string
	fmt.Scan(&n, &direction)
	direction = strings.ToUpper(strings.TrimSpace(strings.Trim(direction, "\r\n")))
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	//steps := make([]string, 0)
	var steps []string
	if direction == "R" {
		steps = rotateRight(matrix)
	} else if direction == "L" {
		steps = rotateLeft(matrix)
	} else {
		fmt.Println(0)
		return
	}
	//log.Println(matrix)

	fmt.Println(len(steps))
	for _, step := range steps {
		fmt.Print(step)
	}
}
