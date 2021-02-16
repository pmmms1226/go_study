package main

import "fmt"

func main() {
	//익명함수를 사용할 경우 함수를 변수에 할당해서 사용가능
	// 함수 안에서 함수를 선언, 정의 가능 --> 이 때 외부 함수에 선언 된 변수에 접근 가능한 함수
	//함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근하기 위한 스냅샷
	//누적 카운터 같은 행위를 할 수 있다.
	//클로저를 정확하게 이해하고 사용해야 한다 --> 자원낭비 심해짐

	//예제1
	multiply := func(x, y int) int {
		return x * y
	} //익명 함수

	r1 := multiply(5, 10)
	fmt.Println(r1)

	//예제
	m, n := 5, 10
	sum := func(c int) int { // 변수가 캡처 됨 (m , n)
		return m + n + c // 지역변수를 사용할 수 있다ㅓ.
	}
	r2 := sum(10)
	fmt.Println(r2)

}
