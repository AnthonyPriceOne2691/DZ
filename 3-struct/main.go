package main

import (
	"fmt"
	"jsoncli/api"
	"jsoncli/bins"
	"jsoncli/config"
	"jsoncli/storage"

	"github.com/joho/godotenv"
)

func main() {
	var store storage.Storage = storage.FileStorage{}

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	cfg := config.NewConfig()
	client := api.NewClient(cfg)

	fmt.Println("API Key from config:", client.APIKey)

	bin := bins.NewBin("1", "MyBin", false)
	list := bins.BinList{Bins: []bins.Bin{bin}}

	err = store.Save("bins.json", list)
	if err != nil {
		panic(err)
	}

	loadedList, err := store.Load("bins.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(loadedList)
}
