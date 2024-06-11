package main

import (
	"encoding/json"
	"fmt"
)

type info struct {
	Marcet []struct {
		Name  string
		Gramm int
		Price int
	}
}

type Name struct {
	Name string
}

func main() {
	text := `{"marcet" : [{"name" : "Ролы Горячие", "gramm" : 1000, "price" : 1129 }, {"name" : "Запеченые", "gramm" : 1450, "price" : 2000 }, {"name" : "Холодные", "gramm" : 2000, "price" : 2300 }]}`
	var infos info
	json.Unmarshal([]byte(text), &infos)
	fmt.Println(infos.Marcet[0].Name)

	for i := range infos.Marcet {
		fmt.Println(i, infos.Marcet[i].Name, infos.Marcet[i].Gramm, infos.Marcet[i].Price)
	}
}
