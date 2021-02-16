package main

import "fmt"

func main() {
	//맵 조회 및 순회

	map1 := map[string]string{
		"hello":  "world",
		"hello2": "world2",
	}

	for k, v := range map1 {
		fmt.Println(k, v)
	}

}
