package main

import "fmt"

func main() {
	//배열
	// 배열은 용량, 길이가 항상 같다
	// 배열 vs 슬라이스 차이점 중요
	// 길이고정 vs 길이가변
	//값 타입 vs 참조 타입
	//복사 전달 vs 참조 값 전달
	// 전체 비교연산자 사용가능 vs 비교 연산자 불가
	// 대부분 슬라이스를 사용한다

	//예제 1 - 배열
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	//var arr2 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr4)
	fmt.Println(arr1)

	fmt.Printf("%-5T %d %v\n", arr1, len(arr1), arr1)
}
