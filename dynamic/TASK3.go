package dynamic

import (
	"fmt"
	"slices"
)

func MinPizzaCost(s, m, l, cs, cm, cl, x int) int {
	chest := Chest{
		wt:  []int{s, m, l},
		val: []int{cs, cm, cl},
	}
	fmt.Println(chest)
	tmp := slices.Max(chest.val)
	n := 3 // количество типов пицц
	matrix := make([][]int, n+1)
	for i := range matrix {
		for j := 0; j <= x; j++ {
			matrix[i] = append(matrix[i], x*tmp)
		}

	}
	for item := 1; item <= n; item++ { // переберём все предметы из сундука
		for capacity := 1; capacity <= x; capacity++ {
			// всё ниже — о рюкзаке вместимостью capacity
			mincostWithoutCurrent := matrix[item-1][capacity]   //  минимальная стоимость предыдущих предметов
			weight := chest.wt[item-1] * (x / chest.wt[item-1]) //  для хранения максимальной стоимости, если положим только
			// текущий предмет
			if x%chest.wt[item-1] != 0 {
				weight += chest.wt[item-1]
			}
			mincostWithCurrent := chest.val[item-1] * (weight / chest.wt[item-1])
			weightOfCurrent := chest.wt[item-1] // сантиметры текущего
			if capacity < weightOfCurrent {
				remainingCapacity := (-1) * (capacity - weightOfCurrent)
				w := (x - remainingCapacity) / weightOfCurrent
				//fmt.Println(remainingCapacity)
				mincostWithCurrent = matrix[item-1][remainingCapacity] + w*chest.val[item-1]
			}
			//if capacity >= weightOfCurrent {    // проверяем, влезет ли текущий предмет в рюкзак
			//	// если текущий влез, то смотрим, что ещё взять
			//	maxcostWithCurrent = chest.val[item-1]                  // сначала положим текущий предмет
			//	remainingCapacity := capacity - weightOfCurrent         // проверим, осталось ли место
			//	maxcostWithCurrent += matrix[item-1][remainingCapacity] // максимальная стоимость оставшегося места
			//}
			fmt.Println(mincostWithCurrent, mincostWithoutCurrent)
			res := min(mincostWithoutCurrent, mincostWithCurrent)
			matrix[item][capacity] = res
		}
		fmt.Println(matrix[item])
	}
	return matrix[n][x]
}
