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
	// p1:= person{
	// 	First: "Hadiuzzaman",
	// }

	// p2:= person{
	// 	First: "Rony",
	// }

	// xp := []person{p1, p2}
	// bs, err := json.Marshal(xp)

	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("this is json data ", string(bs))

	// xp2 := []person{}

	// err = json.Unmarshal(bs, &xp2)

	// if err != nil{
	// 	log.Panic(err)
	// }
	// fmt.Println("back to go format ", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request){

}

func bar(w http.ResponseWriter, r *http.Request){
	
}