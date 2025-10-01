package storage

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBin(id string, name string, private bool) Bin {
	return Bin{
		ID:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}
}

func Save(filename string, list BinList) error {
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func Load(filename string) (BinList, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return BinList{}, err
	}

	if len(data) == 0 {
		return BinList{}, errors.New("файл пустой")
	}

	var list BinList
	err = json.Unmarshal(data, &list)
	if err != nil {
		return BinList{}, err
	}

	return list, nil
}
