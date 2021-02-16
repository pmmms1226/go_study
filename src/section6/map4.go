package main

import "fmt"

func main() {

	//맵 조회할 경우 주의할 점

	map1 := map[string]int{
		"apple":  15,
		"hello":  115,
		"hello2": 123123,
	}

	v1 := map1["apple"]
	v2 := map1["hello3"]
	v3, ok := map1["hello3"] // nil 인지 확인 가능

	fmt.Println(v1, v2, v3, ok)

	if _, ok := map1["hello3"]; ok {

	}

}
