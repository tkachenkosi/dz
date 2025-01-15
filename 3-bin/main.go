package main

import (
	"fmt"

	// "dzw/3-bin/api"
	// "dzw/3-bin/bins"
	"dzw/3-bin/storage"
)

func main() {
	fmt.Println("Start")

	stor := storage.NewStorage()
	stor.PrintInfo()
}
