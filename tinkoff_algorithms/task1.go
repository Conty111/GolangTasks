/*
1 задание
Ограничение времени
1 секунда
Ограничение памяти
64 МБ
Олег — настоящий герой, чьи школьные будни наполнены заботами, уроками и оценками. Он изо всех сил старается, но, как и любой человек, он время от времени допускает ошибки и получает не самые лучшие оценки.

Сегодня Олег стоит перед особенным испытанием — ему предстоит показать своим родителям свои оценки.
Родители попросили показать ему все его оценки за какие-то последовательные  дней.
Оценки представляют собой последовательность целых чисел от  до  включительно — по одной оценке на каждый день.
Олег хочет выбрать такой непрерывный отрезок своих оценок, чтобы в этом отрезке не было оценок  и , а количество оценок  было максимальным.
Помогите Олегу найти этот особенный момент, когда его школьный свет преобладает над тьмой, и его оценки сияют наиболее ярко!

Формат входных данных

Первая строка содержит одно натуральное число  — количество оценок (1≤n≤10^5). Вторая строка содержит  целых чисел — по оценке  за каждый день (2≤m≤5).

Формат выходных данных

Выведите количество пятерок в выбранном Олегом отрезке, удовлетворяющем всем условиям. Если такого отрезка не существует, выведите −1.

Примеры данных
Пример 1
9
5 5 4 5 4 5 4 5 4
9
5 5 4 5 4 5 4 5 4
4
4
Пример 2
8
3 4 4 4 4 5 4 5
8
3 4 4 4 4 5 4 5
2
2
Пример 3
10
5 5 5 5 5 3 5 5 5 5
10
5 5 5 5 5 3 5 5 5 5
-1

*/

package tinkoff_algorithms

import (
	"fmt"
	"math/rand"
	"time"
)

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Task1() {
	var n int
	fmt.Scan(&n)
	arr := make([]uint8, n)
	var ans int = -1
	var res, first int
	var days uint = 1
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		//log.Println(ans, res, first, days)
		if arr[i] == 5 {
			res++
		}
		if days == 7 {
			ans = max1(ans, res)
			days--
			if arr[first] == 5 {
				res--
			}
			first++
		}
		if arr[i] < 4 {
			first = i + 1
			res = 0
			days = 0
		}
		days++
	}

	fmt.Println(ans)
}

func TaskForTest1(n int, arr []int) int {
	//var n int
	//fmt.Scan(&n)
	//arr := make([]uint8, n)
	var ans int = -1
	var res, first int
	var days uint = 1
	for i := 0; i < n; i++ {
		//fmt.Scan(&arr[i])
		//log.Println(ans, res, first, days)
		if arr[i] == 5 {
			res++
		}
		if days == 7 {
			ans = max1(ans, res)
			days--
			if arr[first] == 5 {
				res--
			}
			first++
		}
		if arr[i] < 4 {
			first = i + 1
			res = 0
			days = 0
		}
		days++
	}
	return ans
}

func generateRandomArray(n int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(1) + 2 // генерация случайных оценок от 2 до 5 включительно
	}
	return arr
}

func StressTest() {
	numTests := 10 // количество тестов для проведения стресс-тестирования
	for i := 0; i < numTests; i++ {
		n := rand.Intn(100000) + 1 // генерация случайного значения n от 1 до 100000
		arr := generateRandomArray(n)
		start := 9500
		for j := start; j < start+8; j++ {
			arr[j] = rand.Intn(1) + 4
		}
		t1 := time.Now()
		result := TaskForTest1(n, arr)
		total := time.Now().Sub(t1)
		fmt.Printf("Test %d - Result: %d; Time: %f\n", i+1, result, total.Seconds())
	}
}
