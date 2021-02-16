package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	//구조체
	//go의 필드들의 집합체 또는 컨테이너
	// 필드는 갖지만 메소드는 갖지 않는다.
	// 객체지향 방식을 지원 --> 리시버를 통해 메소드랑 연결
	// 상속, 객체, 클래스 개념 없음

	//예제1
	kim := Account{number: "245-234234", balance: 10000000, interest: 0.015}
	lee := Account{number: "1234-1234123", balance: 12000000}
	park := Account{number: "1231234123-123123123123", interest: 0.025}
	cho := Account{"12389-0128397y", 150000, 0.03}

	fmt.Println(lee)
	fmt.Println(park)
	fmt.Println(cho)
	kim.interest = 10
	fmt.Println(kim)

	//예졔2
	fmt.Println(int(kim.Calculate()))

}
