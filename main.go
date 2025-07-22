package main

import (
	"fmt"
	"math"
)

const bmiPower = 2

func main() {
	fmt.Println("__ Калькулятор индекса массы тела __")
	for {
		userHeight, userWeight := getUserInput()
		BMI := calculateBMI(userHeight, userWeight)
		outputResult(BMI)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(bmi float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", bmi)
	fmt.Println(result)
	switch {
	case bmi < 18.5:
		fmt.Println("У вас сильный дефицит массы тела")
	case bmi < 25: 
		fmt.Println("У вас нормальная масса тела")
	case bmi < 30: 
		fmt.Println("У вас избыточная масса тела")
	default: fmt.Println("У вас степень ожирения")
	}
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

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Вы хотите сделать ещё один расчёт? (y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}