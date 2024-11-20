package main

import "fmt"

func main() {
	transactions := [6]int{1, 2, 3, 4, 5, 6}
	transactionPartial := transactions[1:5]         // [2, 3, 4, 5] transactionPartial
	transactionNewPartial := transactionPartial[:1] // [2] transactionNewPartial
	transactionNewPartial[0] = 30                   // [30] transactionNewPartial, [30 3 4 5] transactionPartial, [1 30 3 4 5 6] transaction

	transactionNewPartial = transactionNewPartial[0:4] // [30 3 4 5] transactionNewPartial, [30 3 4 5] transactionPartial, [1 30 3 4 5 6] transaction

	fmt.Println(transactions)                                           // [1 30 3 4 5 6]
	fmt.Println(transactionPartial)                                     // [30 3 4 5]
	fmt.Println(transactionNewPartial)                                  // [30 3 4]
	fmt.Println(len(transactionPartial), cap(transactionPartial))       // len=4 cap=5
	fmt.Println(len(transactionNewPartial), cap(transactionNewPartial)) // len=3 cap=5
}
