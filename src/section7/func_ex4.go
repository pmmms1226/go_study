package main

import "fmt"

func main() {
	//익명 함수
	// 즉시 실행

	func() {
		fmt.Println("hello world")
	}() //--> 바로 실행함

	//예제 2
	msg := "hbello world"

	func(m string) {
		fmt.Println(m)
	}(msg)
	//예제 3

	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	r := func(x, y int) int {
		return x * y
	} //(10,20) --> 즉시 실행)
	fmt.Println(r(10, 20))

}
