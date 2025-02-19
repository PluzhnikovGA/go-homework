package bins

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

type Bin struct {
	Id string
	Private bool
	CreatedAt time.Time
	Name string
}

func (bin *Bin) NewBin(reader *bufio.Reader) {
	fmt.Println("__ Create Bin __")

	fmt.Print("Enter a new name of Bin: ")
	for {
		input, err := reader.ReadString('\n')

		if err != nil || input == "\n" {
			fmt.Println("Error! Enter a new name of Bin: ")
			continue
		}

		bin.Name = input[:len(input)-1]
		break
	}

	bin.CreatedAt = time.Now()
	bin.Private = false
	bin.Id = fmt.Sprintf("bin-%d", time.Now().Unix())
}

type BinList struct {
	Bins []Bin
}

func (binList *BinList) NewListBin(reader *bufio.Reader) {
	fmt.Println("__ Create BinList __")
	var bin Bin

	Main:
	for {
		bin.NewBin(reader)

		binList.Bins = append(binList.Bins, bin)

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