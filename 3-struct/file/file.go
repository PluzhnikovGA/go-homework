package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(filename string) (data []byte, err error) {
	return os.ReadFile(filename);
}

func WriteFile(content []byte, name string) (bool) {
	file, err := os.Create(name)

	if err != nil {
		fmt.Println(err)
		return false
	}

	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		fmt.Println("Error! Failed to write data")
		return false
	}

	return true
}

func ValidationJSONExtension(name string) (isJson bool) {
	return filepath.Ext(name) == ".json"
}
