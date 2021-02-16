package main

import "fmt"

func multiply(x, y int) (r int) {
	r = x * y
	return r
}
func sum(x, y int) (r int) {
	r = x + y
	return r
}

func main() {

	//함수를 변수에 할당하여 사용
	//예제1(슬라이스에 할당)
	f := []func(int, int) int{multiply, sum}

	a := f[0](10, 10)
	fmt.Println(a)

	//예제2 (변수에 할당)
	//
	//var f1 func(int, int) int = multiply
	//f2 := multiply

	//예제3 (맵에 할당)

	m := map[string]func(int, int) int{
		"mul_func": multiply,
		"sum_func": sum,
	}

}
