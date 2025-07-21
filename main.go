package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight := 1.8
	userWeight := 100.0
	bmi := userWeight / math.Pow(userHeight, 2)
	fmt.Print(bmi)
}
