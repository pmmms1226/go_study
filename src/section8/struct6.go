package main

import "fmt"

type Car struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

type spec struct {
	length int "정장"
	height int "전고"
	width  int "전축"
}

func main() {
	//중첩 구조체
	//예제1
	car1 := Car{
		"520d",
		"silver",
		"bmw",
		spec{4, 5, 1},
	}
	fmt.Println(car1)
}
