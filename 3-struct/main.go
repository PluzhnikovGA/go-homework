package main

import (
	"3-struct/bins"
	"3-struct/storage"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	binList := bins.NewBinList(reader, storage.NewStorageDb("data.json"))

	for {
		fmt.Print("Do you want to add new bin? (Y/n): ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Something went wrong.")
			continue
		}

		switch strings.ToUpper(strings.TrimSpace(input)) {
		case "Y", "":
			binList.AddBin(reader)
			data, err := binList.Store.Read()

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(string(data))
		case "N":
			fmt.Println("Good bye!")
		default:
			fmt.Println("You entered an incorrect value.")
		}
	}
}
