package main

import "fmt"

func main() {
	var a int = 3
	var b int // 기본값 0이 될것이다.
	var c = 4 // type을 생략하면 4에 맞추어 int형으로 초기화될것이다.
	d := 5    // var, type모두 생략. 이 경우도 int형 초기화됨.
	var e = "Hello"
	f := 3.14

	fmt.Println(a, b, c, d, e, f)
}
