package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(filename string) (data []byte, err error) {
	return os.ReadFile(filename);
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		fmt.Println("Error! Failed to write data")
		return
	}

	fmt.Println("Recording completed successfully!")
}

func ValidationJSONExtension(name string) (isJson bool) {
	return filepath.Ext(name) == ".json"
}
