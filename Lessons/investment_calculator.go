package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("expectedReturnRate: ")
	outputText("expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("years: ")
	outputText("years: ")
	fmt.Scan(&years)

	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f", futureRealValue)
	// Outputs information
	// fmt.Println("Future Value:", futureValue)

	//Rounded Places !!!
	//fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for Inflation): %.1f", futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
	fmt.Print(formattedFV, formattedFRV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}
