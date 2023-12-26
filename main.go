package main

import (
	"fmt"
	"github.com/Conty111/GolangTasks/sync_patterns"
	"log"
)

func main() {
	go sync_patterns.Start()
	//log.Println(sync_patterns.Compare("test", "good"))
	log.Println(sync_patterns.Average([]string{"test", "good", "bad", "test", "test"}))
	log.Println(sync_patterns.BestStudents([]string{"test", "good", "bad", "test", "test"}))
}

func t2() {
	var t int
	fmt.Scan(&t)
	for ti := 0; ti < t; ti++ {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}
		if n < 3 {
			fmt.Println("Yes")
			continue
		}
		mx := make([][]int, n)
		can := true
		for {
			min_val := 1000000001
			min_idx := -1
			max_val := 0
			max_idx := -1
			for idx, val := range arr {
				if min_val > val && val > 0 {
					min_val = val
					min_idx = idx
				}
				if max_val < val {
					max_val = val
					max_idx = idx
				}
			}
			//fmt.Println(arr, min_val, max_val)
			if min_idx == max_idx {
				if min_idx == -1 {
					break
				}
				for idx, val := range arr {
					if max_val == val && idx != min_idx {
						max_val = val
						max_idx = idx
						break
					}
				}
				if min_idx == max_idx {
					can = false
					break
				}
			}
			arr[min_idx] -= 1
			arr[max_idx] -= 1
			mx[min_idx] = append(mx[min_idx], max_idx)
			mx[max_idx] = append(mx[max_idx], min_idx)
		}
		if !can {
			fmt.Println("No")
			continue
		}
		used := make([]bool, n)
		var dfs func(v int)
		dfs = func(v int) {
			//fmt.Println(v, used)
			used[v] = true
			for _, u := range mx[v] {
				if !used[u] {
					dfs(u)
				}
			}
		}
		dfs(0)
		isAll := true
		for _, val := range used {
			if !val {
				isAll = false
				break
			}
		}
		if !isAll {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
	}
}

func t1() {
	var t int
	fmt.Scan(&t)
	right_lett := make(map[rune]int, 7)
	for _, l := range "TINKOFF" {
		right_lett[l] += 1
	}
	for i := 0; i < t; i++ {
		var str string
		fmt.Scanln(&str)
		if len(str) != 7 {
			fmt.Println("No")
			continue
		}
		letters := make(map[rune]int, 7)
		for _, let := range str {
			letters[let] += 1
			val, ok := right_lett[let]
			if !ok || val < letters[let] {
				fmt.Println("No")
				break
			}
		}
		fmt.Println("Yes")
	}
}
