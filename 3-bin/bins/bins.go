package bins

import (
	// "encoding/json"
	// "fmt"
	"time"
	// "dzw/3-bin/file"
)

type Bin struct {
	Id       string    `json:"id"`
	Private  bool      `json:"private"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"createAt"`
}

type BinList struct {
	Bins     []Bin     `json:"bins"`
	UpdateAt time.Time `json:"updateAt"`
}

func newBin(id, name string, private bool) (*Bin, error) {
	return &Bin{
		Id:       id,
		Private:  private,
		Name:     name,
		CreateAt: time.Now(),
	}, nil
}

func NewBinList() *BinList {
	return &BinList{
		Bins:     []Bin{},
		UpdateAt: time.Now(),
	}

}

/*
func NewBinList() *BinList {

		file, err := file.ReadFile(stor.file)
		if err != nil {
			return &bins.BinList{
				Bins:     []bins.Bin{},
				UpdateAt: time.Now(),
			}
		}

	var bins BinList
	err := json.Unmarshal(file, &bins)
	if err != nil {
		fmt.Println(err.Error())
		return &BinList{
			Bins:     []Bin{},
			UpdateAt: time.Now(),
		}
	}

	return &bins
}

*/
