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
	fmt.Println(transactions)
}

func scanTransaction() float64 {
	fmt.Print("Введите вашу транзакцию (n для выхода): ")
	var transaction float64
	fmt.Scan(&transaction)
	return transaction
}