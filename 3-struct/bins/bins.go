package bins

import (
	"3-struct/storage"
	"bufio"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Bin struct {
	Id string `json:"id"`
	Private bool `json:"private"`
	CreatedAt time.Time `json:"createAt"`
	Name string `json:"name"`
}

func newBin(reader *bufio.Reader) (*Bin){
	var bin Bin
	fmt.Println("__ Create Bin __")

	fmt.Print("Enter a new name of Bin: ")
	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error! Enter a new name of Bin: ")
			continue
		}

		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Error! Enter a valid name of Bin: ")
			continue
		}

		bin.Name = input
		break
	}

	bin.CreatedAt = time.Now()
	bin.Private = false
	bin.Id = fmt.Sprintf("bin-%d", time.Now().Unix())

	return &bin
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func (binList *BinList) AddBin(reader *bufio.Reader) {
	bin := newBin(reader)

	binList.Bins = append(binList.Bins, *bin)

	storage.SaveBinList(json.MarshalIndent(binList, "", "  "))
}

func NewBinList(reader *bufio.Reader) (*BinList) {
	fmt.Println("__ Create BinList __")

	data, err := storage.ReadBins()

	if err != nil {
		fmt.Println("Error! Something went wrong, please try again.")
		return nil
	}

	var binList BinList
	err = json.Unmarshal(data, &binList)

	if err != nil {
		fmt.Printf("Couldn't parse data from the file")
		return &BinList{
			Bins: []Bin{},
		}
	}

	return &binList
}
