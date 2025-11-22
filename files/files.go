package files

import (
	"os"
)

func ReadFile(name string) ([]byte, error) {
    data, err := os.ReadFile(name)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func WriteFile(content []byte, name string) error {
    return os.WriteFile(name, content, 0600)  
}