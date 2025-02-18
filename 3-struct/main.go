package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Bin struct {
	id string
	private bool
	createdAt time.Time
	name string
}

func (bin *Bin) newBin(reader *bufio.Reader) {
	fmt.Println("__ Create Bin __")

	fmt.Print("Enter a new name of Bin: ")
	for {
		input, err := reader.ReadString('\n')
	
		if err != nil || input == "\n" {
			fmt.Println("Error! Enter a new name of Bin: ")
			continue
		}

		bin.name = input[:len(input)-1]
		break
	}

	bin.createdAt = time.Now()
	bin.private = false
	bin.id = fmt.Sprintf("bin-%d", time.Now().Unix())
}

type BinList struct {
	bins []Bin
}

func (binList *BinList) newListBin(reader *bufio.Reader) {
	fmt.Println("__ Create BinList __")
	var bin Bin

	Main:
	for {
		bin.newBin(reader)

		binList.bins = append(binList.bins, bin)

		fmt.Print("Do you want create new Bin yet [N/y]: ")

		Second:
		for {
			input, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("Error! Do you want create new Bin yet [N/y]: ")
				continue
			}

			input = strings.TrimSpace(input)
			switch input {
			case "N", "", "n":
				break Main
			case "Y", "y":
				break Second
			default:
				fmt.Print("Invalid input. Please enter 'N' or 'Y': ")
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var binList BinList
	binList.newListBin(reader)

	fmt.Println("\nCreated bins:")
	for _, bin := range binList.bins {
		fmt.Printf("ID: %s, Name: %s, CreatedAt: %s Private: %t \n", bin.id, bin.name, bin.createdAt, bin.private)
	}
}
