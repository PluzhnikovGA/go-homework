package storage

import (
	"3-struct/file"
	"fmt"
	"os"
)

type Storage interface {
	Read() (data []byte, err error)
	Save(data []byte, err error) (bool)
}

type StorageDb struct {
	filename string
}

func NewStorageDb(name string) *StorageDb {
	return &StorageDb{
		filename: name,
	}
}

func (db *StorageDb) Read() (data []byte, err error) {
	isJSON := file.ValidationJSONExtension(db.filename)
	if !isJSON {
		fmt.Println("Your didn't specify a JSON file, please ty again.")
		return
	}

	if _, err := os.Stat(db.filename); os.IsNotExist(err) {
		defaultData := []byte(`{"bins":[]}`)

		err = os.WriteFile(db.filename, defaultData, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}
		return defaultData, nil
	}

	return file.ReadFile(db.filename)
}

func (db *StorageDb) Save(data []byte, err error) (bool) {
	if err != nil {
		fmt.Println("Failed to convert to JSON")
		return false
	}

	return file.WriteFile(data, db.filename)
}
