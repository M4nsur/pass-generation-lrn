package account

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte) error
}

type Storage struct {
	Accounts  []account         `json:"accounts"`
	UpdatedAt time.Time         `json:"updatedAt"`    
	db        Db 
}

func CreateStorage(db Db) *Storage {
	data, err := db.Read()
	if err != nil {
		return &Storage{
			Accounts:  []account{},
			UpdatedAt: time.Now(),
			db:        db,  
		}
	}
	var storage Storage
	err = json.Unmarshal(data, &storage)
	if err != nil {
		fmt.Println(err.Error())
		return &Storage{
			Accounts:  []account{},
			UpdatedAt: time.Now(),
			db:        db,
		}
	}

	storage.db = db 
	return &storage
}
func (storage *Storage) FindByUrl() {
	fmt.Println("Введите url для поиска")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	url := strings.TrimSpace(scanner.Text())

	for _, acc := range storage.Accounts {
		if acc.Url == url {
			fmt.Println("Найден аккаунт:", acc)
			return
		}
	}
	fmt.Println("Аккаунт не найден")
}

func (storage *Storage) DeleteByUrl() error {
	fmt.Println("Введите url для удаления")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	url := strings.TrimSpace(scanner.Text()) 
	found := false
	for i, acc := range storage.Accounts {
		if acc.Url == url {
			storage.Accounts = append(storage.Accounts[:i], storage.Accounts[i+1:]...)
			fmt.Println("Аккаунт удален:", acc)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("аккаунт с URL %s не найден", url)
	}

	return storage.Save()
}



func (storage *Storage) Save() error {
	storage.UpdatedAt = time.Now()

	data, err := json.Marshal(storage)
	if err != nil {
		return fmt.Errorf("ошибка при сериализации: %w", err)
	}

	if err := storage.db.Write(data); err != nil { 
		return fmt.Errorf("ошибка при записи: %w", err)
	}

	return nil
}


