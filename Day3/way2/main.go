package main

import (
	"fmt"
	"math"
)

// Complete the solve function below.
func solve(meal_cost float64, tip_percent int32, tax_percent int32) {

	tip1 := float64(tip_percent)
	tax1 := float64(tax_percent)
	tip := (meal_cost * (tip1 / 100))
	tax := (meal_cost * (tax1 / 100))

	total := meal_cost + tip + tax
	total = math.Round(total)

	fmt.Println(int64(total))

}

func main() {

	var meal_cost float64
	var tip_percent, tax_percent int32

	fmt.Println("Enter meal_cost,meal_tip,meal_tax")

	fmt.Scanln(&meal_cost)
	fmt.Scanln(&tip_percent)
	fmt.Scanln(&tax_percent)

	solve(meal_cost, tip_percent, tax_percent)
}
