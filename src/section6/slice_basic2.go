package main

import "fmt"

func main() {
	//슬라이스(참조타입 증명)

	//예제1(배열)
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1
	arr2[0] = 7

	fmt.Println(arr2[0])

	//예제2 (슬라이스)
	slice1 := []int{1, 2, 3}
	var slice2 []int
	slice2 = slice1
	slice2[0] = 7

	fmt.Println(slice1)

	//예외 상황

	slice3 := make([]int, 50, 100) // 길이만큼만 0으로 초기화 된
	fmt.Println(slice3[4])
	fmt.Println(slice3[50])

}
