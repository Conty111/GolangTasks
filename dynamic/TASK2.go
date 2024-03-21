package dynamic

import "sort"

type Chest struct {
	val []int
	wt  []int
}

func Knapsack(chest *Chest, maxWeight int) (int, []int) {
	n := len(chest.val) // количество драгоценностей
	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, maxWeight+1)
	}
	matrix_items := make([][][]int, n+1)
	for i := range matrix_items {
		matrix_items[i] = make([][]int, maxWeight+1)
	}

	for item := 1; item <= n; item++ { // переберём все предметы из сундука
		for capacity := 1; capacity <= maxWeight; capacity++ {
			// всё ниже — о рюкзаке вместимостью capacity
			maxcostWithoutCurrent := matrix[item-1][capacity] //  максимальная стоимость предыдущих предметов
			itemsWithoutCurrent := matrix_items[item-1][capacity]
			itemsWithCurrent := []int{item - 1}
			maxcostWithCurrent := 0 //  для хранения максимальной стоимости, если положим текущий предмет

			weightOfCurrent := chest.wt[item-1] // масса текущего
			if capacity >= weightOfCurrent {    // проверяем, влезет ли текущий предмет в рюкзак
				// если текущий влез, то смотрим, что ещё взять
				maxcostWithCurrent = chest.val[item-1]                  // сначала положим текущий предмет
				remainingCapacity := capacity - weightOfCurrent         // проверим, осталось ли место
				maxcostWithCurrent += matrix[item-1][remainingCapacity] // максимальная стоимость оставшегося места
				itemsWithCurrent = append(itemsWithCurrent, matrix_items[item-1][remainingCapacity]...)
			}
			res := max(maxcostWithoutCurrent, maxcostWithCurrent)
			var res_items []int
			if res == 0 {
				res_items = []int{}
			} else if res == maxcostWithoutCurrent {
				res_items = itemsWithoutCurrent
			} else {
				res_items = itemsWithCurrent
			}
			matrix[item][capacity] = res
			matrix_items[item][capacity] = res_items
		}
	}
	sort.Ints(matrix_items[n][maxWeight])
	return matrix[n][maxWeight], matrix_items[n][maxWeight]
}
