package bank

import (
	"errors"
	"fmt"
)

type Account struct {
	ID      int
	Holder  string
	Balance int
	History []string
}

type Transaction struct {
	Description string
	Amount      int
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
	a.History = append(a.History, fmt.Sprintf("Пополнение: +%d", amount))
}

func (a *Account) Withdraw(amount int) error {
	a.Balance -= amount
	if a.Balance < amount {
		return errors.New("Недостаточно денег")
	}
	a.History = append(a.History, fmt.Sprintf("Снятие: -%d", amount))
	return nil
}

func (a Account) PrintInfo() {
	fmt.Printf("ID: %d Владелец: %s Баланс: %d\n", a.ID, a.Holder, a.Balance)
}

func Transfer(from *Account, to *Account, amount int) error {
	fmt.Printf("Попытка перевода %d руб. от %s к %s\n", amount, from.Holder, to.Holder)

	err := from.Withdraw(amount)
	if err != nil {
		return err
	}
	to.Deposit(amount)
	return nil
}
