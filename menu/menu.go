package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/m4nsur/pass-generation-lrn/account"
)

func RunMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	storage := account.CreateAccountStorage("data.json")

	for {
		showMenu()
		scanner.Scan()
		choiceStr := strings.TrimSpace(scanner.Text())
		choice, err := strconv.Atoi(choiceStr)
		
		if err != nil {
			fmt.Println("Неверная команда")
			continue
		}
		
		switch choice {
		case 1:
			err := account.NewAccount(scanner, storage)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 2:
			storage.FindAccount()
		case 3:
			storage.DeleteAccount("data.json")
		case 4:
			fmt.Println("Программа завершена")
			return
		default:
			fmt.Println("Неверная команда")
		}
	}
}

func showMenu() {
	fmt.Println("\nМеню:")
	fmt.Println("1.Создать аккаунт")
	fmt.Println("2.Найти аккаунт")
	fmt.Println("3.Удалить аккаунт")
	fmt.Println("4.Выход")
	fmt.Print("Выберите пункт меню: ")
}