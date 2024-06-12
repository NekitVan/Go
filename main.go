package main

import (
	"encoding/json"
	"fmt"
)

type info struct {
	ttt []struct {
		Name  string
		Date  int
		Price int
	}
}

func rol() {
	text := `{"ttt" : [{"name" :  "Бутер", "date" : 10.06.2024 , "price" : 1200} ,{"name" :  "Солянка", "date" : 13.06.2024 , "price" : 500}, {"name" :  "Пюре", "date" : 09.06.2024 , "price" : 1000}]}`
	var infos info
	json.Unmarshal([]byte(text), &infos)
	fmt.Println(infos.ttt[0].Name)

	for i := range infos.ttt {
		fmt.Println(i, infos.ttt[i].Name, infos.ttt[i].Date, infos.ttt[i].Price)
	}
}

func main() {
	rol()
}
