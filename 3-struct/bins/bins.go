package bins

import (
	"3-struct/storage"
	"encoding/json"
	"fmt"
	"time"
)

type Bin struct {
	Id string `json:"id"`
	Private bool `json:"private"`
	CreatedAt time.Time `json:"createAt"`
	Name string `json:"name"`
}

func newBin(name string) (*Bin){
	var bin Bin

	bin.Name = name
	bin.CreatedAt = time.Now()
	bin.Private = false
	bin.Id = fmt.Sprintf("bin-%d", time.Now().Unix())

	return &bin
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

type BinListWithStore struct {
	BinList
	Store storage.Storage
}

func (binList *BinListWithStore) AddBin(name string) {
	binList.Bins = append(binList.Bins, *newBin(name))

	binList.save(binList.Bins)
}

func (binList *BinListWithStore) Update(id string) {
	for index := range binList.Bins {
		if binList.Bins[index].Id == id {
			binList.Bins[index].CreatedAt = time.Now()
			binList.save(binList.Bins)
			fmt.Printf("Bin with id: %s was updated\n", id)
			return
		}
	}

	fmt.Printf("Bin list don't have bin with id: %s", id)
}

func (binList *BinListWithStore) Delete(id string) {
	for index := range binList.Bins {
		if binList.Bins[index].Id == id {
			binList.Bins = append(binList.Bins[:index], binList.Bins[index+1:]...)
			binList.save(binList.Bins)
			fmt.Printf("Bin with id: %s was deleted", id)
			return
		}
	}

	fmt.Printf("Bin list don't have bin with id: %s", id)
}

func (binList *BinListWithStore) Get(id string) {
	for _, bin := range binList.Bins {
		if bin.Id == id {
			fmt.Printf("id: %s\n", bin.Id)
			fmt.Printf("name: %s\n", bin.Name)
			fmt.Printf("private: %t\n", bin.Private)
			fmt.Printf("createdAt: %s\n", bin.CreatedAt)
			return
		}
	}

	fmt.Printf("Bin list don't have bin with id: %s", id)
}

func (binList *BinListWithStore) GetList() {
	if len(binList.Bins) == 0 {
		fmt.Println("Bin list is empty.")
		return
	}

	for _, bin := range binList.Bins {
		fmt.Printf("id: %s, name: %s\n", bin.Id, bin.Name)
	}
}

func NewBinList(store *storage.StorageDb) (*BinListWithStore) {
	data, err := store.Read()

	if err != nil {
		fmt.Println("Error! Something went wrong, please try again.")
		return nil
	}

	var bins BinList
	err = json.Unmarshal(data, &bins)

	if err != nil {
		fmt.Printf("Couldn't parse data from the file")
		return &BinListWithStore{
			BinList: BinList{
				Bins: []Bin{},},
			Store: store,
		}
	}

	return &BinListWithStore{
		BinList: bins,
		Store: store,
	}
}

func (binList *BinListWithStore) save(bins []Bin) {
	tempBinList := BinList{
		Bins: bins,
	}

	binList.Store.Save(json.MarshalIndent(tempBinList, "", "  "))
}
