package main

import (
	"fmt"
	"cli/bin"
)


func main() {
	b := bin.NewBin("123", true, "2024-06-01T12:00:00Z", "MyBin")
	fmt.Printf("Bin: %+v\n", b)
	bl := bin.NewBinList()
	bl = append(bl, b)
	fmt.Printf("BinList: %+v\n", bl)
}