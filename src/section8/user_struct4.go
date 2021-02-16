package main

import "fmt"

type shoppingBasket struct {
	cnt, price int
}

func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

//원본 수정(참조 형식)
func (b *shoppingBasket) rePurchase(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

//원본 수정 불가
func (b shoppingBasket) rePurchase2(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	//리시버(구조체 메소드) 전달(값, 참조) 형식이 있음
	// 함수는 기본적으로 값 호출 -> 변수의 값이 복사 후 내부 전달(원본 수정X) 맵, 슬라이스 등은 참조 전달
	//리시버도 마찬가지로 포인터를 활용해서 메소드 내에서 원본 수정 가능

	//예제1
	bs1 := shoppingBasket{
		3, 5000,
	}
	fmt.Println("ex1 tot price : ", bs1.purchase())
	bs1.rePurchase(7, 5000)
	fmt.Println("ex1 tot price : ", bs1.purchase()) //원본 수정 --> 포인터로 선언해도 메소드를 호출할때 별도로 해줄건 없다. 알아서 해준다.
	bs1.rePurchase2(10, 0)
	fmt.Println("ex1 tot price : ", bs1.purchase()) //원본 수정 X

}
