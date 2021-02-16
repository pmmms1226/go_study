package main

import "fmt"

func main() {
	//슬라이스 복사
	//make 로 공간을 확보 후 복사를 해야한다.
	//복사 된 슬라이스 값 변경해도 원본에는 영향 없음

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	//slice3 := []int{} --> 복사가 안됨
	copy(slice2, slice1)
	fmt.Println(slice2)
}
