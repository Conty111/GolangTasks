package main

import (
	"fmt"
	"github.com/Conty111/GolangTasks/game_of_life"
	"time"
)

func main() {
	newWorld := game_of_life.NewWorld(10, 10)
	newWorld.Seed()
	for {
		fmt.Print(newWorld.String())
		newWorld.Next()
		time.Sleep(200 * time.Millisecond)
		// специальная последовательность для очистки экрана после каждого шага
		fmt.Print("\033[H\033[2J")
	}
}
