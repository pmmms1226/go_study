package main

import "fmt"

//사용자 정의 타입 == 함수편

type totCost func(int, int) int

func describe(cnt int, price int, fn totCost) {
	fmt.Println(cnt, price, fn(cnt, price))
}

func main() {
	var orderPrice totCost
	orderPrice = func(cnt int, price int) int {
		return (cnt * price)
	}
	describe(10, 20, orderPrice)
}
