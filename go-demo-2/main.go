package main

import "fmt"

func main() {
	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	sum := calculateBalance(transactions)
	fmt.Printf("Ваш баланс: %.2f", sum)
}

func scanTransaction() float64 {
	fmt.Print("Введите вашу транзакцию (n для выхода): ")
	var transaction float64
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(arr[]float64) float64 {
	sum := 0.0
	for _, value := range arr {
		sum += value
	}
	return sum
}