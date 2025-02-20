package storage

import (
	"3-struct/file"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadBins() (data []byte, err error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter name of file for reading: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}

		input = strings.TrimSpace(input)

		isJSON := file.ValidationJSONExtension(input)
		if !isJSON {
			fmt.Println("Your didn't specify a JSON file, please ty again.")
			continue
		}

		data, err := file.ReadFile(input)

		return data, err
	}
}

func SaveBinList(data []byte, err error) {
	if err != nil {
		fmt.Println("Failed to convert to JSON")
		return
	}

	file.WriteFile(data, "data.json")
}
