package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight float64 = 1.8
	var userWeight = 100
	var bmi = float64(userWeight) / math.Pow(userHeight, 2)
	fmt.Print(bmi)
}
