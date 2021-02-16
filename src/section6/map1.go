package main

import "fmt"

func main() {
	//맵 : 해시테이블, 딕셔너리(파이썬), key-value 값
	//레퍼런스 타입(참조 타입)
	//비교 연산자 사용 불가능 (참조 타입이므로)
	//특징: 참조타입(key)로 사용 불가능, 값으로는 모든 타입 사용가능
	//순서가 없음
	//make 및 리터럴로 초기화 가능

	//예제1
	//var map1 map[string]int = make(map[string]int)
	//
	//map2 := make(map[string]int)
	//
	map3 := map[string]int{}
	map3["hello"] = 45
	fmt.Println(map3)

	map4 := map[string]int{
		"hello": 6,
		"world": 10,
	}
	fmt.Println(map4)

	map5 := make(map[string]int, 10)
	map5["hello"] = 1515

}
