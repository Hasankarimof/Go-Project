package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5
	expectedReturnRate, investmentAmount, years := 1000.0, 10.0, 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("years: ")
	fmt.Scan(&years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureRealValue)
}
