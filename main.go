package main

import (
	lock_free_DS "github.com/Conty111/GolangTasks/lock-free_DS"
)

func main() {
	var a1 int32 = 23
	var a2 int32 = 5
	lock_free_DS.AtomicSwap(&a1, &a2)
	println(a1, a2)
}
