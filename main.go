package main

import (
	"encoding/json"
	"fmt"
)

type MenuItem struct {
	Name  string  `json:"name"`
	Date  string  `json:"date"`
	Price float64 `json:"price"`
}

type Info struct {
	TTT []MenuItem `json:"ttt"`
}

func findItemByDate(date string, infos Info) {
	faund := false
	for _, item := range infos.TTT {
		if item.Date == date {
			fmt.Printf("Найден блюдо '%s' с ценой %.2f рублей.", item.Name, item.Price)
			faund = true
			break
		}
	}
	if !faund {
		fmt.Println("Нет тавара в наличии!")
	}

}
func rol() Info {
	text := `{"ttt" : [{"name" :  "Бутер", "date" :"10.06.2024" , "price" : 1200} ,{"name" :  "Солянка", "date" : "13.06.2024" , "price" : 500}, {"name" :  "Пюре", "date" : "09.06.2024" , "price" : 1000}]}`
	var infos Info
	json.Unmarshal([]byte(text), &infos)
	fmt.Println(infos.TTT)

	for i := range infos.TTT {
		fmt.Println(i, infos.TTT[i].Name, infos.TTT[i].Date, infos.TTT[i].Price)
	}
	return infos
}

func main() {
	infos := rol()
	findItemByDate("10.06.2024", infos)
}
