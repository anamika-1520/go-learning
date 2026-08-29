package main

import "fmt"

type BankAccount struct {
    Owner   string
    Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
    a.Balance += amount
}

func (a *BankAccount) Withdraw(amount float64) {
    if amount <= a.Balance {
        a.Balance -= amount
    } else {
        fmt.Println("Insufficient balance")
    }
}

func (a BankAccount) ShowBalance() {
    fmt.Println("Balance:", a.Balance)
}

func main() {

    account := BankAccount{
        Owner:   "Anamika",
        Balance: 1000,
    }

    account.ShowBalance()

    account.Deposit(500)
    account.ShowBalance()

    account.Withdraw(200)
    account.ShowBalance()
}