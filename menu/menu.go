package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/m4nsur/pass-generation-lrn/account"
	"github.com/m4nsur/pass-generation-lrn/files"
)

func RunMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	storage := account.CreateStorage(files.CreateJsonDb("data.json"))

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
			storage.FindByUrl()
		case 3:
			storage.FindByUrl()	
		case 4:
			err := storage.DeleteByUrl()
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 5:
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
	fmt.Println("2.Найти аккаунт по url")
	fmt.Println("3.Найти аккаунт по login")
	fmt.Println("4.Удалить аккаунт")
	fmt.Println("5.Выход")
	fmt.Print("Выберите пункт меню: ")
}



// func promptData[T any](promptSlice []T)  {
// 	for i, item := range promptSlice {
// 		if (i == len(promptSlice) - 1) {
// 			fmt.Printf("%v:", item)
// 		} else {
// 			fmt.Println(item)
// 		}
// 	}
// }