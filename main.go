package main

import (
	"fmt"
	"math"
)

const bmiPower = 2

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
		 	continue
		}
		fmt.Printf("%d\n", i)
	}

	fmt.Println("__ Калькулятор индекса массы тела __")
	userHeight, userWeight := getUserInput()
	BMI := calculateBMI(userHeight, userWeight)
	outputResult(BMI)
	switch {
	case BMI < 18.5:
		fmt.Println("У вас сильный дефицит массы тела")
	case BMI < 25: 
		fmt.Println("У вас нормальная масса тела")
	case BMI < 30: 
		fmt.Println("У вас избыточная масса тела")
	default: fmt.Println("У вас степень ожирения")
	}
}

func outputResult(bmi float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", bmi)
	fmt.Println(result)
}

func calculateBMI(userHeight float64, userWeight float64) float64 {
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