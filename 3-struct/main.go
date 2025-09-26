package main

import (
	"fmt"
	"reflect"
	"time"
)

type Bin struct {
	ID        string 	 `json:"id"`
	Private   bool    	 `json:"private"`
	CreatedAt time.Time  `json:"created_at"`
	Name      string     `json:"name"`
}

type BinList struct {
	Bins []Bin
}

func NewBin(id string, name string, private bool) Bin {
	return Bin{
		ID:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}
}


func main() {
	field, _ := reflect.TypeOf(Bin{}).FieldByName("Private")
	fmt.Println(field.Tag)
}
