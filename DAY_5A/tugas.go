// Masih on progress, masih perlu diubah kedalam bentuk fungsi terpisah
package main

import (
	"fmt"
)

type Goods struct {
	id    int
	name  string
	price int
}

func main() {
	goods := []Goods{
		{id: 1, name: "Benih Lele", price: 50000},
		{id: 2, name: "Pakan Lele Cap Menara", price: 25000},
		{id: 3, name: "Probiotik A", price: 75000},
		{id: 4, name: "Probiotik Nila B", price: 10000},
		{id: 5, name: "Pakan Nila", price: 20000},
		{id: 6, name: "Benih Nila", price: 20000},
		{id: 7, name: "Cupang", price: 5000},
		{id: 8, name: "Benih Nila", price: 30000},
		{id: 9, name: "Benih Cupang", price: 10000},
		{id: 10, name: "Probiotik B", price: 10000},
	}
	for key, element := range goods {
		fmt.Println(key, "-", element)
	}

	// urut untuk total < 100000
	fmt.Println("urut untuk total < 100000")
	for i := 0; i < len(goods)-1; i++ {
		for j := 0; j < len(goods)-i-1; j++ {
			if goods[j].price > goods[j+1].price {
				goods[j], goods[j+1] = goods[j+1], goods[j]
			}
		}
	}
	itemList := []Goods{}
	sum := 0
	for _, element := range goods {
		if sum < 100000 {
			itemList = append(itemList, element)
			sum += element.price
		} else {
			break
		}

	}
	fmt.Println(itemList)

	// cari yg 10000
	fmt.Println("Mencari yg 10000")
	for _, element := range goods {
		if element.price == 10000 {
			fmt.Println(element.name, "-", element.price)
		}
	}
	//
	goods2 := goods[0]

	// Cari paling murah
	fmt.Println("Cari paling murah")
	for i := 0; i < len(goods); i++ {
		if goods[i].price < goods2.price {
			goods2 = goods[i]
		}
	}
	fmt.Println("Paling Murah", goods2)

	//Cari paling mahal
	fmt.Println("Cari paling mahal")
	for i := 0; i < len(goods); i++ {
		if goods[i].price > goods2.price {
			goods2 = goods[i]
		}
	}
	fmt.Println("Paling Mahal", goods2)
}
