package main

import (
	"fmt"
	"sort"
)

func main() {
	//슬라이스 추출 및 정렬
	//slice[i:j] i ~ j-1 까지 추출
	//slice[:] 처음부터 끝까지 추출

	//예제1
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println(slice1[:])

	//예제 2
	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println(sort.IntsAreSorted(slice2))
}
