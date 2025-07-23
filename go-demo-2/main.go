package main

import "fmt"

func main() {
	tr1 := []int{1, 2, 3}
	tr2 := []int{4, 5, 6}
	tr1 = append(tr1, tr2...)
	fmt.Println(tr1)
	
	for _, value := range tr1 {
		fmt.Println(value)
	}

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