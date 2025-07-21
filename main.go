package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ Калькулятор индекса массы тела __")
	userHeight, userWeight := getUserInput()
	BMI := calculateBMI(userHeight, userWeight)
	outputResult(BMI)
}

func outputResult(bmi float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", bmi)
	fmt.Print(result)
}

func calculateBMI(userHeight float64, userWeight float64) float64 {
	const bmiPower = 2
	BMI := userWeight / math.Pow(userHeight / 100, bmiPower)
	return BMI
}

func getUserInput() (float64, float64)  {
	var userHeight float64
	var userWeight float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}