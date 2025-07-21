package main

import (
	"fmt"
	"math"
)

func main() {
	const bmiPower = 2
	userHeight := 1.8
	userWeight := 100.0
	bmi := userWeight / math.Pow(userHeight, bmiPower)
	fmt.Print(bmi)
}
