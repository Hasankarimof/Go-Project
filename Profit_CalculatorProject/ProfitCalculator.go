package main

import (
	"fmt"
)

func main() {

	// Create variables
	var revenue int
	var expenses int
	var taxRate float64

	// Ask the user for input
	fmt.Print("\nrevenue: ")
	fmt.Scan(&revenue)
	fmt.Print("\nexprenses: ")
	fmt.Scan(&expenses)
	fmt.Print("\ntax rate: ")
	fmt.Scan(&taxRate)

	// Calculate EBT
	EBT := revenue - expenses

	// Calculate Profit
	Profit := float64(EBT) * (1 - taxRate/100)

	// Calculate ratio
	Ratio := float64(EBT) / Profit

	// Output the result
	fmt.Println(EBT)
	fmt.Println(Profit)
	fmt.Println(Ratio)

}
