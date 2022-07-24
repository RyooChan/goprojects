package main

import "fmt"

type User struct {
	Name string
	Id   string
	Age  int
}

type VIPUser struct {
	Userinfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"류찬", "chan", 26}
	vip := VIPUser{
		User{"김찬", "kimChan", 27},
		3,
		500,
	}

	fmt.Printf("유저:%s ID:%s 나이:%d \n", user.Name, user.Id, user.Age)
	fmt.Printf("VIP 유저:%s ID:%s 나이:%d VIP레벨:%d VIP유지비용:%d만원\n",
		vip.Userinfo.Name,
		vip.Userinfo.Id,
		vip.Userinfo.Age,
		vip.VIPLevel,
		vip.Price,
	)
}
