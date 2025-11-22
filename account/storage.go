package account

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)


type AccountsStorage struct {
	Accounts []account
	UpdatedAt time.Time `json:"updatedAt"`
}


func (storage *AccountsStorage) FindAccount () {
	println("Введите url для поиска")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()             
    url := scanner.Text()

	for _, acc := range storage.Accounts {
		isFound := strings.Contains(acc.Url, url)
		if isFound {
			fmt.Println(acc)
		}
	}

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

