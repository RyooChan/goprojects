package main

import "fmt"

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	var house House
	house.Address = "경기도 군포시"
	house.Size = 52
	house.Price = 10
	house.Category = "아파트"

	fmt.Println(house)
	fmt.Printf("주소:%s 사이즈:%d평 가격:%f억 종류:%s\n", house.Address, house.Size, house.Price, house.Category)
}
