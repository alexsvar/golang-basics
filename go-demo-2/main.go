package main

import "fmt"

func main() {
	transactions := []int{10, 20, 30}
	temp := transactions
	transactions = append(transactions, 50)
	newTransactions := append(transactions, 100, 200);

	fmt.Println(temp)
	fmt.Println(transactions)
	fmt.Println(newTransactions)
}