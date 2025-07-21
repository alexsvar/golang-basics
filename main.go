package main

import (
	"fmt"
	"math"
)

func main() {
	const bmiPower = 2
	var userHeight float64
	var userWeight float64
	fmt.Println("__ Калькулятор индекса массы тела __")
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	bmi := userWeight / math.Pow(userHeight / 100, bmiPower) 
	fmt.Printf("Ваш индекс массы тела: %.0f", bmi)
}
