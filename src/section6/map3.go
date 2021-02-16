package main

import "fmt"

func main() {

	//맵 값 변경 및 삭제
	map1 := map[string]string{
		"hello":  "world",
		"hello2": "world2",
	}

	fmt.Println(map1)

	map1["hello3"] = "world3"

	fmt.Println(map1)

	//삭제

	delete(map1, "hello3")

	fmt.Println(map1)

}
