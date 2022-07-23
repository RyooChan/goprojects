package main

import "fmt"

func main() {
	var a int16 = 3456
	var b int8 = int8(a)

	fmt.Println(a, b) // 이러면 3456, -128 이렇게 나온다.
}
