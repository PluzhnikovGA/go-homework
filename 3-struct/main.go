package main

import (
	"3-struct/storage"
	"fmt"
)

func main() {
	data, _ := storage.ReadBins()

	fmt.Print(data)
	// reader := bufio.NewReader(os.Stdin)

	// var binList bins.BinList
	// binList.NewListBin(reader)

	// fmt.Println("\nCreated bins:")
	// for _, bin := range binList.Bins {
	// 	fmt.Printf("ID: %s, Name: %s, CreatedAt: %s Private: %t \n", bin.Id, bin.Name, bin.CreatedAt, bin.Private)
	// }
}
