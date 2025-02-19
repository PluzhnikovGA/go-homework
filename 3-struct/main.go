package main

import (
	"3-struct/bins"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var binList bins.BinList
	binList.NewListBin(reader)

	fmt.Println("\nCreated bins:")
	for _, bin := range binList.Bins {
		fmt.Printf("ID: %s, Name: %s, CreatedAt: %s Private: %t \n", bin.Id, bin.Name, bin.CreatedAt, bin.Private)
	}
}
