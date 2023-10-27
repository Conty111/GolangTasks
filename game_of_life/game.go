package game_of_life

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

const brownSquare = "\xF0\x9F\x9F\xAB "
const greenSquare = "\xF0\x9F\x9F\xA9 "
const testSquare = "\xF0\x9F\x9F\xA2 "

func NewWorld(height, width int) *World {
	world := &World{
		Height: height,
		Width:  width,
	}
	world.Cells = make([][]bool, height)
	for i := 0; i < height; i++ {
		world.Cells[i] = make([]bool, width)
	}
	return world
}

func (w *World) Seed() {
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			if rand.Intn(5) == 1 {
				w.Cells[i][j] = true
			}
		}
	}
}

func (w *World) LoadState(filename string) error {
	inF, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	defer inF.Close()
	if err != nil {
		return err
	}
	sc := bufio.NewScanner(inF)
	var i int
	var row_len int
	var res [][]bool
	for sc.Scan() {
		res_row := sc.Text()
		if i == 0 {
			row_len = len(res_row)
		}
		if row_len != len(res_row) {
			return fmt.Errorf("Width is not right")
		}
		res = append(res, make([]bool, row_len))
		for j := 0; j < row_len; j++ {
			if res_row[j] == 49 {
				res[i][j] = true
			} else {
				res[i][j] = false
			}
		}
		i++
	}
	w.Height = i
	w.Width = row_len
	w.Cells = res
	return nil
}

func (w *World) SaveState(filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	defer func(f *os.File) { _ = f.Close() }(f)
	if err != nil {
		return err
	}
	for i := 0; i < w.Height; i++ {
		var row string
		for j := 0; j < w.Width; j++ {
			if w.Cells[i][j] {
				row += "1"
			} else {
				row += "0"
			}
		}
		if i != w.Height-1 {
			row += "\n"
		}
		_, err := f.Write([]byte(row))
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *World) String() string {
	var res string
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			if w.Cells[i][j] {
				res += greenSquare
			} else {
				res += brownSquare
			}
		}
		res += "\n"
	}
	return res
}

func (w *World) Next() {
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			w.Cells[i][j] = w.NextState(j, i)
		}
	}
}

func (w *World) Neighbors(x, y int) int {
	var res, y1, x1 int
	for i := y - 1; i <= y+1; i++ {
		if i == -1 {
			y1 = w.Height - 1
		} else if i == w.Height {
			y1 = 0
		} else {
			y1 = i
		}
		for j := x - 1; j <= x+1; j++ {
			if !(j == x && i == y) {
				if j == -1 {
					x1 = w.Width - 1
				} else if j == w.Width {
					x1 = 0
				} else {
					x1 = j
				}
				if w.Cells[y1][x1] {
					res += 1
				}
			}
		}
	}
	return res
}

func (w *World) NextState(x, y int) bool {
	current_state := w.Cells[y][x]
	neighboardsLive := w.Neighbors(x, y)
	if neighboardsLive == 3 {
		return true
	}
	if current_state && neighboardsLive == 2 {
		return true
	}
	return false
}
