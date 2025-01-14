package main

import (
	"fmt"
	"time"

	"dzw/3-bin/api"
	"dzw/3-bin/bins"
	"dzw/3-bin/file"
	"dzw/3-bin/storage"
)

type Bin struct {
	id       string
	private  bool
	createAt time.Time
	name     string
}

type BinList struct {
	id string
	Bin
}

func newBin(id, name string, private bool) (*Bin, error) {
	return &Bin{
		id:       id,
		private:  private,
		createAt: time.Now(),
		name:     name,
	}, nil
}

func newBinList(id, name string, private bool) (*BinList, error) {
	return &BinList{
		id: id,
		Bin: Bin{
			private:  private,
			createAt: time.Now(),
			name:     name,
		},
	}, nil
}

func main() {
	fmt.Println("Info message")

}
