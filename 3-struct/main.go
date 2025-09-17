package main

import (
	"time"
)

type Bin struct {
	ID        string   
	Private   bool      
	CreatedAt time.Time 
	Name      string    
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

}
