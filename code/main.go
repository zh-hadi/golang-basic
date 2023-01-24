package main

import (
	"encoding/json"
	"log"
	"fmt"
	
)

type person struct {
	First string
}


func main(){
	p1:= person{
		First: "Hadiuzzaman",
	}

	p2:= person{
		First: "Rony",
	}

	xp := []person{p1, p2}
	bs, err := json.Marshal(xp)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(bs))
}