package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.

	// Read and save an integer, double, and String to your variables.

	// Print the sum of both integer variables on a new line.

	// Print the sum of the double variables on a new line.

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.

	var i2 uint64
	var d2 float64
	fmt.Scanln(&i2)
	fmt.Scanln(&d2)
	scanner.Scan()
	s2 := scanner.Text()

	fmt.Println(i + i2)
	fmt.Printf("%.1f\n", d+d2)

	// or use Sprintf()
	//NewString :=fmt.Sprintf("%s%s",s,s1)
	fmt.Println(s + s2)

}
