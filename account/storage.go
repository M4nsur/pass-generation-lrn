package account

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)


type AccountsStorage struct {
	Accounts []account
	UpdatedAt time.Time `json:"updatedAt"`
}

func CreateAccountStorage () *AccountsStorage {
	data, err := os.ReadFile("data.json")
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
