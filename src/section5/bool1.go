package main

import "fmt"

func main() {
	//Bool (Boolean): 참과 거짓
	// 조건부 논리 연산자랑 주로 사용: !, || , &&
	// 암묵적 형변환X : 0, Nil => false가 아님 --> 자바나 c 처럼 1,0 으로 사용할 수 없다.

	var b1 bool = true
	var b2 bool = false
	b3 := true
	fmt.Println("b1: ", b1)
	fmt.Println("b2: ", b2)
	fmt.Println("b3: ", b3)

}
