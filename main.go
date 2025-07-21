package main

import (
	"fmt"
	"math"
)

func main() {
	const bmiPower = 2
	var userHeight float64
	var userWeight float64
	fmt.Print("__ Калькулятор индекса массы тела __\n")
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	bmi := userWeight / math.Pow(userHeight, bmiPower)
	fmt.Print("Ваш индекс массы тела: ")
	fmt.Print(bmi)
}
