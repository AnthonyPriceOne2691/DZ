package main

import (
	"fmt"
	"jsoncli/bins"
	"jsoncli/storage"
)

func main() {
	var store storage.Storage = storage.FileStorage{}

	bin := bins.NewBin("1", "MyBin", false)
	list := bins.BinList{Bins: []bins.Bin{bin}}

	err := store.Save("bins.json", list)
	if err != nil {
		panic(err)
	}

	loadedList, err := store.Load("bins.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(loadedList)
}
