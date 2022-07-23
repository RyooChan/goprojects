package main

import "fmt"

func main() {
	a := 3 // int -> int64
	var b float64 = 3.5

	var c int = int(b) // float64를 int형태로 바꾼다.
	d := float64(a) * b

	var e int64 = 7 // int64
	f := a * int(e) // int(int64)랑 int64으로 같은것같은데.. 그래도 변환 없이는 안된다. 왜냐면 type이 다르기 때문.

	fmt.Println(a, b, c, d, e, f)
}
