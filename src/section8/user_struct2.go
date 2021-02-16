package main

import "fmt"

//사용자 정의 타입
type cnt int

func testConverT(i int) {
	fmt.Println("ex3: default", i)
}

func testConverD(i cnt) {
	fmt.Println("ex3: custom", i)
}

func main() {
	//기본 자료형 사용자 정의 타입
	//예제1

	a := cnt(5)

	//예제 2
	var b cnt = 15
	fmt.Println(a, b)

	//예제 3
	//testConverT(b) //int는 int를 무조건 넣어야한다.
	testConverT(int(b)) //이렇게 강제변환 해야지 사용이 가능하다.

}
