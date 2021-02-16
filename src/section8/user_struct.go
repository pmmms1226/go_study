package main

import "fmt"

type Car struct {
	//멤버 변수
	name  string
	color string
	price int64
	tax   int64
}

//일반 메소드
func Price(c Car) int64 {
	return c.price + c.tax
}

// 구조체 <-> 메소드 바인드
func (c Car) Price() int64 {
	return c.tax + c.price
}

func main() {
	//GO -> 객체 지향 타입을 구조체로 정의한다. (클ㄹ래스, 상속 개념 없음)
	//객체 지향 -> 클래스(속성: 멤버변수, 기능(상태: 메소드)) : 코드의 재사용성, 관리가 용이, 신뢰성이 높은 프로그래밍
	//고는 객체지향 언어일까? 맞음
	// Go는 전형적인 객체지향의 특징을 가지고 있지 않지만, 인터페이스 -> 다형성 지원, 구조체를 클래스 형태의 코딩 가능
	// 객체지향의 기본 개념 -> Go에서 포함하고 있다. -> 객체 지향 프로그래밍 언어
	// 구조체를 사용해야함
	//상태 , 메소드를 분리해서 정의(결합성 없음)
	// 사용자 정의 타입: 구조체, 인터페이스, 기본타입, 함수
	// 구조체 != 클래스 --> 구조체와 메소드는 분리된 개념
	// 구조체 <-> 메소드 바인딩 필

	//예제1
	bmw := Car{
		name:  "520d",
		price: 50000000,
		color: "withe",
		tax:   5000000,
	}
	benz := Car{
		name:  "220d",
		price: 60000000,
		color: "withe",
		tax:   6000000,
	}
	fmt.Println("ex1: ", bmw, &bmw)
	fmt.Println("ex1: ", benz, &benz)

	fmt.Println("ex2: ", Price(bmw))

	fmt.Println("ex3: ", bmw.Price())

}
