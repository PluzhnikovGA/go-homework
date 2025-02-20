package file

import (
	"fmt"
	"os"
	"path/filepath"
)

const EXT = ".json"

func ReadFile(name string) (data []byte, err error) {
	return os.ReadFile(name);
}

func ValidationJSONExtension(name string) (isJson bool) {
	return filepath.Ext(name) == EXT
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
