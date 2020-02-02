package main

import (
	"fmt"
)

func main() {
	var n int32
	fmt.Println("Enter Any Nos")
	fmt.Scanf("%d", &n)
	for i := 1; i < 11; i++ {
		result := n * int32(i)

		fmt.Println(n, "x", i, "=", result)
	}
}
