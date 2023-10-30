package main

import (
	"fmt"
	"github.com/Conty111/GolangTasks/OOP"
)

func main() {
	//world := game_of_life.NewWorld(10, 30)
	//world.Seed()
	//for {
	//	fmt.Println(world.String())
	//	world.Next()
	//	time.Sleep(50 * time.Millisecond)
	//	fmt.Print("\033[H\033[2J")
	//}
	var test OOP.Vehicle
	test = OOP.Motorcycle{Speed: 12.4, Type: "SomeType", FuelType: "Some1"}
	fmt.Println(test)
}
