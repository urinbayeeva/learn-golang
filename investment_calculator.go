package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate  = 6.5
	var investmentAmount float64
	var years float64
	expectedReturnedRate := 5.5
    
	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return rate: ")
	fmt.Scan(&expectedReturnedRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)


	futureValue := float64(investmentAmount) * math.Pow(1 + expectedReturnedRate / 100, (years)) 
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}