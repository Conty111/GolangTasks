package main

import (
	"fmt"

	"example.com/v2/dynamic"
)

func main() {
	nums := []int{3, 9, 10, 1, 30, 40}
	res := dynamic.MaxExpressionValue(nums)
	fmt.Println(res)
}
