package account

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/m4nsur/pass-generation-lrn/files"
)


type AccountsStorage struct {
	Accounts []account
	UpdatedAt time.Time `json:"updatedAt"`
}


func (storage *AccountsStorage) FindAccount() {
	fmt.Println("Введите url для поиска")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	url := scanner.Text()

	for _, acc := range storage.Accounts {
		if acc.Url == url {
			fmt.Println("Найден аккаунт:", acc)
			return
		}
	}
	fmt.Println("Аккаунт не найден")
}

func (storage *AccountsStorage) DeleteAccount(storageName string) error {
	fmt.Println("Введите url для удаления")
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()             
    url := scanner.Text()

	for i, acc := range storage.Accounts {
	    if acc.Url == url {  
	        storage.Accounts = append(storage.Accounts[:i], storage.Accounts[i+1:]...)
	        fmt.Println("Аккаунт удален:", acc)
	        break  
	    }
	}

	data, err := ToBytes(storage)
	if err != nil {
		fmt.Println("Ошибка при сериализации:", err)
		return err
	}

	err = files.WriteFile(data, storageName)
	if err != nil {
		fmt.Println("Ошибка при сохранении:", err)
		return err
	}

	return nil
}




func CreateAccountStorage (storageName string) *AccountsStorage {
	data, err := os.ReadFile(storageName)
	if (err != nil) {
		return &AccountsStorage{
			Accounts: []account{},
			UpdatedAt: time.Now(),
		}
	}

	var storage AccountsStorage
	err = json.Unmarshal(data, &storage)
	if (err != nil) {
		fmt.Println(err.Error())
	}
	return &storage
}

func ToBytes(acc *AccountsStorage) ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

