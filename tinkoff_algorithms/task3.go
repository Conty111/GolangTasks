package tinkoff_algorithms

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Dir struct {
	Name   string
	Childs map[string]*Dir
}

func Task3() {
	var n int
	fmt.Scan(&n)
	var root *Dir
	for i := 0; i < n; i++ {
		path, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		path = strings.TrimSpace(path)
		curDirs := strings.Split(path, "/")
		prevDir := root
		for i, dir := range curDirs {
			if i == 0 {
				if root == nil {
					root = &Dir{
						//Parent: nil,
						Name:   dir,
						Childs: make(map[string]*Dir),
					}
					prevDir = root
				}
			} else if _, ok := prevDir.Childs[dir]; !ok {
				newDir := &Dir{Name: dir, Childs: make(map[string]*Dir)}
				prevDir.Childs[dir] = newDir
				prevDir = newDir
			} else if _, ok = prevDir.Childs[dir]; ok {
				prevDir = prevDir.Childs[dir]
			}
		}
	}
	printDir(0, root)
}

func printDir(otstup uint, d *Dir) {
	//log.Println(d.Name)
	fmt.Printf("%s%s\n", strings.Repeat(" ", int(otstup)), d.Name)
	keys := make([]string, 0, len(d.Childs))

	// Заполняем слайс ключами из карты
	for key := range d.Childs {
		keys = append(keys, key)
	}
	// Сортируем ключи
	sort.Strings(keys)
	for _, key := range keys {
		printDir(otstup+2, d.Childs[key])
	}
}
