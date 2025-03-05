package bins

import (
	"3-struct/api"
	"3-struct/storage"
	"encoding/json"
	"fmt"
	"time"
)

type Bin struct {
	Private bool `json:"private"`
	CreatedAt time.Time `json:"createAt"`
	Name string `json:"name"`
	Id string `json:"id"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

type BinListWithStore struct {
	BinList
	Store storage.Storage
}

func NewBin(resp *api.ApiBinResp) (*Bin){
	var bin Bin

	parsedTime := ParsedTime(resp.Metadata.CreatedAt)

	if parsedTime.IsZero() {
		fmt.Println("error parsing date")
		return  &Bin{}
	}

	bin.Name = resp.Record.Name
	bin.CreatedAt = parsedTime
	bin.Private = resp.Metadata.Private
	bin.Id = resp.Metadata.Id

	return &bin
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

func (binList *BinListWithStore) AddBin(resp *api.ApiBinResp) (bool) {
	binList.Bins = append(binList.Bins, *NewBin(resp))

	return binList.save(binList.Bins)
}

func (binList *BinListWithStore) Delete(id string) (error) {
	for index := range binList.Bins {
		if binList.Bins[index].Id == id {
			binList.Bins = append(binList.Bins[:index], binList.Bins[index+1:]...)
			if binList.save(binList.Bins) {
				return nil
			} else {
				return fmt.Errorf("something went wrong, try again")
			}
		}
	}

	return fmt.Errorf("bin list don't have bin with id: %s", id)
}

func (binList *BinListWithStore) GetBin(id string) (*Bin, error) {
	for _, bin := range binList.Bins {
		if bin.Id == id {
			return &bin, nil
		}
	}

	return &Bin{}, fmt.Errorf("Bin list don't have bin with id: %s", id)
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

func ParsedTime(createdAt string) (time.Time){
	parsedTime, err := time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return  time.Time{}
	}

	return parsedTime
}

func (binList *BinListWithStore) save(bins []Bin) (bool) {
	return binList.Store.Save(json.MarshalIndent(getTempleBinList(bins), "", "  "))
}

func getTempleBinList(bins []Bin) *BinList {
	return &BinList{
		Bins: bins,
	}
}
