package main

import "fmt"

// В цикле спрашиваем ввод трансакций: -10, 10, 40.5,
// Добавлять в массив трансакций
// Вывести массив
// Вывести сумму баланса в консоль

func main() {
	transactions := make([]float64, 0, 10)
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	balance := calculateBalance(transactions)
	fmt.Printf("Ваш баланс: %.2f", balance)
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Введите трансакцию (n для выхода): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}
