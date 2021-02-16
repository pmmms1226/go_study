package main

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	//다양한 선언 방법
	//&struct, struct : &struct 포인터를 받아오고 역참조를 또 하기때문에 속도는 조금 느리다.
	//꼭 포인터를 사용해야할때 --> 인터페이스 메소드를 선언만 해둔 후 --> 오버라이딩 해서 메서드에 포인터 리시버를 사용할 경우 반드시 &struct를 사용해야한다.

	//선언방법1
	var kim *Account = new(Account)
	kim.number = "12341234"
	kim.balance = 1000000
	kim.interest = 0.015

	//선언방법2
	hong := &Account{number: "12341234", balance: 123012300, interest: 0.04}

	//선언방법3
	lee := new(Account)

}
