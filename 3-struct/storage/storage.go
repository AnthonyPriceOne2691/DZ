package storage

import (
	"encoding/json"
	"errors"
	"os"

	"jsoncli/bins"
)

type Storage interface {
	Save(filename string, list bins.BinList) error
	Load(filename string) (bins.BinList, error)
}

type FileStorage struct{}

func (fs FileStorage) Save(filename string, list bins.BinList) error {
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (fs FileStorage) Load(filename string) (bins.BinList, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return bins.BinList{}, err
	}

	if len(data) == 0 {
		return bins.BinList{}, errors.New("файл пустой")
	}

	var list bins.BinList
	err = json.Unmarshal(data, &list)
	if err != nil {
		return bins.BinList{}, err
	}

	return list, nil
}
