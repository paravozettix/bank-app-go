package main

import (
	"fmt"
	"mybank/bank"
)

func main() {
	lastID := 1002
	accounts := make(map[int]*bank.Account)
	accounts[1001] = &bank.Account{1001, "Ivan", 1000, []string{}}
	accounts[1002] = &bank.Account{1002, "Elena", 500, []string{}}

	for {
		fmt.Println("\n--- МЕНЮ БАНКА ---")
		fmt.Println("1. Показать все счета")
		fmt.Println("2. Перевести деньги")
		fmt.Println("3. Создать новый счет")
		fmt.Println("4. Пополнить баланс счета")
		fmt.Println("5. История транзакций")
		fmt.Println("6. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			for _, acc := range accounts {
				acc.PrintInfo()
			}
			continue
		case 2:
			var idFrom, idTo, amount int

			fmt.Println("Введите ID отправителя: ")
			fmt.Scan(&idFrom)
			sender, ok1 := accounts[idFrom]
			if !ok1 {
				fmt.Println("Отправитель не найден")
				continue
			}

			fmt.Println("Введите ID получателя: ")
			fmt.Scan(&idTo)
			receiver, ok2 := accounts[idTo]
			if !ok2 {
				fmt.Println("Получатель не найден")
				continue
			}

			fmt.Println("Введите сумму для перевода: ")
			fmt.Scan(&amount)

			err := bank.Transfer(sender, receiver, amount)
			if err != nil {
				fmt.Println("Ошибка перевода:", err)
			} else {
				fmt.Println("Перевод выполнен!")
			}

		case 3:
			var name string
			fmt.Print("Введите имя владельца: ")
			fmt.Scan(&name)
			lastID++

			newAcc := &bank.Account{lastID, name, 0, []string{}}
			accounts[lastID] = newAcc

			fmt.Printf("Ваш счет успешно создан. Ваш ID: %d\n", lastID)

		case 4:
			var id, amount int
			fmt.Print("Введите ID счёт для пополнения: ")
			fmt.Scan(&id)
			acc, ok := accounts[id]
			if !ok {
				fmt.Printf("Счёт ID: %d не был найден", id)
				continue
			}

			fmt.Print("Введите сумму для пополнения: ")
			fmt.Scan(&amount)

			acc.Deposit(amount)

		case 5:
			var id int
			fmt.Print("Введите ID для просмотра истории: ")
			fmt.Scan(&id)

			acc, ok := accounts[id]
			if !ok {
				fmt.Printf("Счёт ID: %d не был найден", id)
				continue
			}

			for _, record := range acc.History {
				fmt.Println("-- " + record)
			}

		case 6:
			return
		}
	}
}
