package main

import (
	"fmt"
	"os"
)

func main() {

	for {

		// Create variables
		// var revenue float64
		// var expenses float64
		// var taxRate float64

		// Ask the user for input
		// fmt.Print("\nrevenue: ")
		revenue := EnterText("\nRevenue: ")
		// fmt.Scan(&revenue)

		// fmt.Print("\nexprenses: ")
		expenses := EnterText("\nExpenses: ")
		// fmt.Scan(&expenses)

		// fmt.Print("\ntax rate: ")
		taxRate := EnterText("\nTaxRate : ")
		// fmt.Scan(&taxRate)

		if revenue <= 0 {
			println("Invalid amount. Must be greater than 0.")
			continue
		} else if expenses <= 0 {
			println("Invalid amount. Must be greater than 0.")
			continue
		} else if taxRate <= 0 {
			println("Invalid amount. Must be greater than 0.")
			continue
		}

		// // Calculate EBT
		// EBT := revenue - expenses

		// // Calculate Profit
		// Profit := float64(EBT) * (1 - taxRate/100)

		// // Calculate ratio
		// Ratio := float64(EBT) / Profit

		EBT, Profit, Ratio := outCalc(revenue, expenses, taxRate)

		// Output the result
		fmt.Printf("%.0f\n", EBT)
		fmt.Printf("%.1f\n", Profit)
		fmt.Printf("%.2f\n", Ratio)

		// Storing Results
		StoreResults(EBT, Profit, Ratio)

		// Stopping Loop
		break
	}

}

func EnterText(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func outCalc(revenue, expenses, taxRate float64) (float64, float64, float64) {

	var EBT = revenue - expenses
	var Profit = EBT * (1 - taxRate/100)
	var Ratio = EBT / Profit

	return EBT, Profit, Ratio
}

func StoreResults(EBT, Profit, Ratio float64) {
	StoreSheet := fmt.Sprintf("EBT:  %.1f\nProfit:  %.1f\nRatio:  %.3f\n", EBT, Profit, Ratio)
	os.WriteFile("CalculatedResults", []byte(StoreSheet), 0644)
}
