package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var no int
	fmt.Scanln(&no)

	for j := 0; j < no; j++ {

		var str string
		fmt.Scanln(&str)
		var s1, s2 string
		for i := 0; i < len(str); i++ {

			if i%2 == 0 {
				s1 = s1 + string(str[i])
			} else {
				s2 = s2 + string(str[i])
			}
		}
		fmt.Println(s1, " ", s2)

	}
}
