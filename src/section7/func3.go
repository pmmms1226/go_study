package main

import "fmt"

//다중값 반환
func multiply(x int, y int) (int, int) {
	return x * 10, y * 10
}

func main() {

	a, b := multiply(10, 5)
	fmt.Println(a, b)

}
