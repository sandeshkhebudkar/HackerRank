package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var name string
	//also use scan method but is only read only one word
	// fmt.Scanln(&name)
	//using reader and scanner
	/* 	reader := bufio.NewReader(os.Stdin)
	   	text, err := reader.ReadString('\n')
	*/
	//# Read using Scanner
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Hello, World.")
	fmt.Println(name)
}
