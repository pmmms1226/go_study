package main

import "fmt"

func rptc(n *int) {
	*n = 77
}
func vptc(n int) {
	n = 77
}

func main() {

	//i := 7
	//p := &i

	//포인터의 값 전달
	// 함수 , 메서드 호출 시에 매개변수 값을 복사 전달 -->
	//크기가 큰 배열 전달을 포인터로 할때는 시스템의 큰 이득

	a := 10
	b := 10

	rptc(&a)
	vptc(b)

	fmt.Println(a, b)
}
