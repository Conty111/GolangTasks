/*
Банковский счет
Напишите программу для управления банковским счетом. Создайте структуру Account с приватными полями balance (баланс) и owner (владелец). Реализуйте методы для установки баланса и получения баланса, а также методы для внесения и снятия денег с счета. Убедитесь, что баланс не может быть отрицательным.

Примечания
Код программы должен содержать описание струкрутры Account:
type Account struct { owner string balance float64 }
*/

package OOP

import "fmt"

type Account struct {
	owner   string
	balance float64
}

func NewAccount(owner string, balance float64) *Account {
	return &Account{
		owner:   owner,
		balance: balance,
	}
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) SetBalance(money float64) error {
	if money < 0 {
		return fmt.Errorf("Net stolko deneg")
	}
	a.balance = money
	return nil
}

func (a *Account) Deposit(money float64) {
	a.balance += money
}

func (a *Account) Withdraw(how_much float64) error {
	if a.balance-how_much < 0 {
		return fmt.Errorf("Net stolko deneg")
	}
	a.balance -= how_much
	return nil
}
