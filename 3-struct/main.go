package main

import (
	"3-struct/bins"
	"3-struct/storage"
	"flag"
	"fmt"
	"os"
)

func main() {
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
		fmt.Println("Error! You need to write file name.")
		os.Exit(1)
	}
	binList := bins.NewBinList(storage.NewStorageDb(*file))

	switch {
	case *create:
		requireValue(name, "bin name")
		binList.AddBin(*name)
	case *update:
		requireValue(id, "bin id")
		binList.Update(*id)
	case *delete:
		requireValue(id, "bin id")
		binList.Delete(*id)
	case *get:
		requireValue(id, "bin id")
		binList.Get(*id)
	case *list:
		binList.GetList()
	default:
		exitWithError(fmt.Sprintln("Error! You need to write a command."))
	}
}

func requireValue(value *string, name string) {
	if *value == "" {
		exitWithError(fmt.Sprintf("Error! You need to write %s.", name))
	}
}

func exitWithError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
