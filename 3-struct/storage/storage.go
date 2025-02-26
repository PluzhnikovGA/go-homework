package storage

import (
	"3-struct/file"
	"fmt"
)

type Storage interface {
	Read() (data []byte, err error)
	Save(data []byte, err error)
}

type StorageDb struct {
	filename string
}

func NewStorageDb(name string) *StorageDb {
	return &StorageDb{
		filename: name,
	}
}

func (db StorageDb) Read() (data []byte, err error) {
	isJSON := file.ValidationJSONExtension(db.filename)
	if !isJSON {
		fmt.Println("Your didn't specify a JSON file, please ty again.")
		return
	}

	return file.ReadFile(db.filename)
}

func (db StorageDb) Save(data []byte, err error) {
	if err != nil {
		fmt.Println("Failed to convert to JSON")
		return
	}

	file.WriteFile(data, db.filename)
}
