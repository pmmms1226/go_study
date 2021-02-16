package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a, b int) {
	fmt.Println(a + b)
}
func main() {
	//함수(콜백) 전달 --> 함수 전달 가능, 참조전달 , 값 전달
	//예제1
	sum(100, add)

}
