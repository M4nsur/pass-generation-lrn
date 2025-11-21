package menu

import (
	"fmt"

	"github.com/m4nsur/pass-generation-lrn/account"
	"github.com/m4nsur/pass-generation-lrn/files"
)

func RunMenu() {
		var choice int8
		for  {
			showMenu()
			fmt.Scan(&choice)
			switch choice {
			case 1:
				myAcc, err:= account.NewAccount()
				if err != nil {
					fmt.Println(err)
					break
				}
				files.WriteFile(myAcc, "data.json")
			case 2:
				files.ReadFile("data.json")
			case 3:
				files.ReadFile("data.json")
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