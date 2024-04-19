package tinkoff_algorithms

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func Task5() {
	var n int
	fmt.Scan(&n)
	forest := make([][]string, n)
	dp := make([][]uint, n)
	for i := 0; i < n; i++ {
		cells, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		forest[i] = strings.Split(strings.TrimSpace(cells), "")
		dp[i] = make([]uint, 3)
	}
	var maxMushrooms uint
	for i := 0; i < 3; i++ {
		if forest[0][i] == "C" {
			dp[0][i]++
		}
	}
	for i := 1; i < n; i++ {
		isBlocked := true
		for j := 0; j < 3; j++ {
			if forest[i][j] != "W" {
				isBlocked = false
				var left, center, right uint
				if forest[i-1][1] != "W" {
					center = dp[i-1][1]
				}
				if j > 0 && forest[i-1][2] != "W" {
					right = dp[i-1][2]
				}
				if j < 2 && forest[i-1][0] != "W" {
					left = dp[i-1][0]
				}
				if forest[i][j] == "C" {
					dp[i][j]++
				}
				dp[i][j] += max(left, max(center, right))
				maxMushrooms = max(dp[i][j], maxMushrooms)
			}
		}
		if isBlocked {
			break
		}
	}
	fmt.Println(maxMushrooms)
}

//func Task5ForTest(n int, forest [][]string) int {
//	//forest := make([][]string, n)
//	dp := make([][]uint, n)
//	for i := 0; i < n; i++ {
//		//forest[i] = strings.Split(strings.TrimSpace(rows[i]), "")
//		dp[i] = make([]uint, 3)
//	}
//	var maxMushrooms uint
//	for i := 0; i < 3; i++ {
//		if forest[0][i] == "C" {
//			dp[0][i]++
//		}
//	}
//	for i := 1; i < n; i++ {
//		isBlocked := true
//		for j := 0; j < 3; j++ {
//			if forest[i][j] != "W" {
//				isBlocked = false
//				var left, center, right uint
//				if forest[i-1][1] != "W" {
//					center = dp[i-1][1]
//				}
//				if j > 0 && forest[i-1][2] != "W" {
//					right = dp[i-1][2]
//				}
//				if j < 2 && forest[i-1][0] != "W" {
//					left = dp[i-1][0]
//				}
//				if forest[i][j] == "C" {
//					dp[i][j]++
//				}
//				dp[i][j] += max(left, max(center, right))
//				maxMushrooms = max(dp[i][j], maxMushrooms)
//			}
//		}
//		if isBlocked {
//			break
//		}
//	}
//	log.Println(dp)
//	return int(maxMushrooms)
//}
//
//func generateForest(h int) [][]string {
//	rand.Seed(uint64(time.Now().Nanosecond()))
//	forest := make([][]string, h)
//	symbols := []string{".", "C", "W"}
//	for i := 0; i < h; i++ {
//		forest[i] = []string{symbols[rand.Intn(2)], symbols[rand.Intn(2)], symbols[rand.Intn(2)]}
//		forest[i][rand.Intn(3)] = "W"
//	}
//	return forest
//}

//func StressTestsTask5() {
//	n := 10000
//	forest := generateForest(n)
//	t1 := time.Now()
//	res := Task5ForTest(n, forest)
//	t2 := time.Now().Sub(t1)
//	log.Printf("Result: %d, Time: %f", res, t2.Seconds())
//}
