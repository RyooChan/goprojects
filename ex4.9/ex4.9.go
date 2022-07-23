package main

import (
	"fmt"
)

func main() {
	nums := [...]int{10, 20, 30, 40, 50} // 크기 5로 초기화

	nums[2] = 300

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
