package main

import (
	"3-struct/api"
	"3-struct/bins"
	"3-struct/config"
	"3-struct/storage"
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
  if err != nil {
		exitWithCode("Error loading .env file", 1)
  }

	create := flag.Bool("create", false, "Create bin")
	update := flag.Bool("update", false, "Update bin")
	delete := flag.Bool("delete", false, "Delete bin")
	get := flag.Bool("get", false, "Get bin")
	list := flag.Bool("list", false, "Bins list")

	file := flag.String("file", "", "File name")
	name := flag.String("name", "", "Bin name")
	id := flag.String("id", "", "Bin id")

	flag.Parse()

	if *file == "" {
		exitWithCode("Error! You need to write file name.", 1)
	}

	binList := bins.NewBinList(storage.NewStorageDb(*file))

	switch {
	case *create:
		createFunc(name, binList)
	case *update:
		updateFunc(id, binList)
	case *delete:
		deleteFunc(id, binList)
	case *get:
		getFunc(id, binList)
	case *list:
		binList.GetList()
	default:
		exitWithCode("Error! You need to write a command.", 1)
	}
}

func requireValue(value *string, name string) {
	if *value == "" {
		exitWithCode(fmt.Sprintf("Error! You need to write %s.", name), 1)
	}
}

func createFunc(name *string, binList *bins.BinListWithStore){
	requireValue(name, "bin name")

	resp, err := api.CreateBin(*config.GetConfig(), *name)

	if err != nil {
		exitWithCode(err.Error(), 1)
	}

	if binList.AddBin(resp) {
		exitWithCode("New bin was added in list", 0)
	} else {
		exitWithCode("New bin wasn't added in list", 1)
	}
}

func getFunc(id *string, binList *bins.BinListWithStore){
	requireValue(id, "bin id")

	fileBin, fileError := binList.GetBin(*id)
	serverBin, serverError := api.GetBin(*config.GetConfig(), *id)

	if fileError != nil && serverError != nil {
		exitWithCode(fmt.Sprintf("You don't have bin with %s", *id), 1)
	} else if fileError != nil {
		exitWithCode(fileError.Error(), 1)
	} else if serverError != nil {
		exitWithCode(serverError.Error(), 1)
	} else if fileBin.Name != serverBin.Record.Name || fileBin.CreatedAt != bins.ParsedTime(serverBin.Metadata.CreatedAt) || fileBin.Private != serverBin.Metadata.Private {
		exitWithCode("File bin have a few different with server bin. Check values.", 1)
	} else {
		printBinInfo(fileBin)
	}
}

func updateFunc(id *string, binList *bins.BinListWithStore) {
	requireValue(id, "bin id")

	bin, err := binList.GetBin(*id)

	if err != nil {
		exitWithCode(err.Error(), 1)
	}

	_, err = api.UpdateBin(*config.GetConfig(), *id, bin.Name)

	if err != nil {
		exitWithCode(err.Error(), 1)
	}

	exitWithCode(fmt.Sprintf("Bin with id %s was updated\n", *id), 0)
}

func deleteFunc(id *string, binList *bins.BinListWithStore) {
	requireValue(id, "bin id")

	err := binList.Delete(*id)

	if err != nil {
		exitWithCode(err.Error(), 1)
	}

	err = api.DeleteBin(*config.GetConfig(), *id)

	if err != nil {
		exitWithCode(err.Error(), 1)
	}

	exitWithCode(fmt.Sprintf("Bin with id %s was deleted\n", *id), 0)
}

func exitWithCode(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}

func printBinInfo(bin *bins.Bin) {
	fmt.Printf("id: %s\n", bin.Id)
	fmt.Printf("name: %s\n", bin.Name)
	fmt.Printf("private: %t\n", bin.Private)
	fmt.Printf("createdAt: %s\n", bin.CreatedAt)
	os.Exit(0)
}
