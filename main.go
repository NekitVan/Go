package main

import (
	"encoding/json"
	"fmt"
)

type Info struct {
	TTT []struct {
		Name  string `json:"name"`
		Date  string `json:"date"`
		Price int    `json:"price"`
	} `json:"ttt"`
}

type Name struct {
	Name string
}

func rol() {
	text := `{"ttt" : [{"name" :  "Бутер", "date" :"10.06.2024" , "price" : 1200} ,{"name" :  "Солянка", "date" : "13.06.2024" , "price" : 500}, {"name" :  "Пюре", "date" : "09.06.2024" , "price" : 1000}]}`
	var infos Info
	json.Unmarshal([]byte(text), &infos)
	fmt.Println(infos.TTT)

	for i := range infos.TTT {
		fmt.Println(i, infos.TTT[i].Name, infos.TTT[i].Date, infos.TTT[i].Price)
	}
}

func main() {
	rol()
}
