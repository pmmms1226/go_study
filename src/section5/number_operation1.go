package main

import (
	"fmt"
	"math"
)

func main() {
	//숫자 연산(산술, 비교)
	//타입이 같아야 가능
	//다른 타입끼리는 반드시 형 변환 후 연산
	// 현 변환이 없을 경우는 예외처리
	//+ - * %  << >> & ^

	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("ex1: ", n1)
	fmt.Println("ex1: ", n2)
	fmt.Println("ex1: ", n3)
	fmt.Println("ex1: ", n4)

}
