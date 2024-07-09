/*
I have an account in the bank and want to deposit and update the balance
*/
package main

import "fmt"

type Account struct {
	holder_name string
	balance     float64
}

func (a *Account) Deposit(deposit_amt float64) {
	fmt.Println("Depositing amount: ", deposit_amt)
	a.balance += deposit_amt
}

func (a *Account) Balance() float64 {
	return a.balance
}

func main() {
	var account Account
	account = Account{holder_name: "Harish", balance: 1000}
	fmt.Println("Opening balance: ", account.Balance())
	account.Deposit(200.00)
	fmt.Println("Current balance is: ", account.Balance())
}

/*
Output:
Opening balance:  1000
Depositing amount:  200
Current Balance is:  1200
*/
