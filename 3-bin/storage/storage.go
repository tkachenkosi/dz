package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"dzw/3-bin/bins"
	"dzw/3-bin/file"
)

const (
	NAME_FILE = "storage.dat"
	NAME_STOR = "Тестовая база"
)

type Storage struct {
	name     string `json:"name"`
	bins     *bins.BinList
	updateAt time.Time `json:"updateAt"`
}

func NewStorage() *Storage {
	data, err := os.ReadFile(NAME_FILE)
	if err != nil {
		fmt.Println(err)
		return &Storage{
			name:     NAME_STOR,
			bins:     bins.NewBinList(),
			updateAt: time.Now(),
		}
	}

	var stor Storage
	err = json.Unmarshal(data, &stor)
	if err != nil {
		fmt.Println(err.Error())
		return &Storage{
			name:     NAME_STOR,
			bins:     bins.NewBinList(),
			updateAt: time.Now(),
		}

	}
	return &stor
}

func (stor *Storage) toBytes() ([]byte, error) {
	file, err := json.Marshal(stor)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (stor *Storage) Save() {
	stor.updateAt = time.Now()
	data, err := stor.toBytes()
	if err != nil {
		fmt.Println(err.Error())
	}
	file.WriteFile(data, NAME_FILE)
}

func (stor *Storage) PrintInfo() {
	fmt.Println("База:", stor.name)
}
