package main

import "fmt"

//매개변수 가변 가능 (여러개 넣을 수 있다)
func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}
	return tot
}

func main() {
	//가변 인자 실습 (매개변수 개수가 동적으로 변할 때 - 정해져 있지 않음)
	fmt.Println(multiply(1, 2, 3, 4, 5, 6))

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	multiply(a...)
}
