package tinkoff_algorithms

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var board [][]string
var used [][]bool
var n int
var endX, endY int
var answer uint

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func Task6() {
	fmt.Scan(&n)
	board = make([][]string, n)
	var startX, startY int
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		cells, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		board[i] = strings.Split(strings.TrimSpace(strings.Trim(cells, "\n")), "")
		for j := range board[i] {
			if board[i][j] == "S" {
				startX, startY = i, j
			} else if board[i][j] == "F" {
				endX, endY = i, j
			}
		}
	}
	used = make([][]bool, n)
	for i := 0; i < n; i++ {
		used[i] = make([]bool, n)
	}
	dfs(startX, startY, "K", 0)
	used = make([][]bool, n)
	for i := 0; i < n; i++ {
		used[i] = make([]bool, n)
	}
	dfs(startX, startY, "G", 0)
	fmt.Println(answer)
}

func dfs(x, y int, figure string, steps uint) {
	used[x][y] = true
	var variants [][]int
	log.Println(x, y, figure, steps)
	if board[x][y] == "K" {
		figure = "K"
	} else if board[x][y] == "G" {
		figure = "G"
	}

	if figure == "K" {
		variants = [][]int{
			{x + 1, y + 1},
			{x + 1, y},
			{x + 1, y - 1},
			{x, y + 1},
			{x, y - 1},
			{x - 1, y + 1},
			{x - 1, y},
			{x - 1, y - 1},
		}
	} else {
		variants = [][]int{
			{x + 2, y + 1},
			{x + 2, y - 1},
			{x + 1, y + 2},
			{x - 1, y + 2},
			{x + 1, y - 2},
			{x - 1, y - 2},
			{x - 2, y + 1},
			{x - 2, y - 1},
		}
	}
	for _, cell := range variants {
		nextX, nextY := cell[0], cell[1]
		if nextX == endX && nextY == endY {
			answer = min(answer, steps+1)
			return
		}
		if nextX < 0 || nextY >= n || nextX >= n || nextY < 0 {
			continue
		}
		if !used[nextX][nextY] {
			dfs(nextX, nextY, figure, steps+1)
		}
	}
}
