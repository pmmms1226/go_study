package main

import "fmt"

func main() {
	//구조체 익명 선언 및 활용
	//예제1 type 구조체명 struct
	car1 := struct {
		name, color string
	}{"520d", "black"}
	fmt.Println(car1)

	//예제2
	cars := []struct{ name, color string }{{"52d", "black"}, {"50d", "blak"}, {"20d", "blak"}}
	for _, c := range cars {
		fmt.Printf("(%s,%s) ----- (%#v)\n", c.name, c.color, c
		)
	}

	//예제3

}
