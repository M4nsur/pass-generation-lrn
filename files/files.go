package files

import (
	"os"
)

type JsonDb struct {
    filename string
}

func (db *JsonDb) Read() ([]byte, error) {
    data, err := os.ReadFile(db.filename)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func (db *JsonDb) Write(content []byte) error {
    return os.WriteFile(db.filename, content, 0600)  
}


func CreateJsonDb (name string) *JsonDb {
    return &JsonDb{
        filename: name,
    }
}
