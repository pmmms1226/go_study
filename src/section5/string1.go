package main

import (
	"fmt"
)

func main() {
	//문자열 "",``
	// Golang: 문자 char 타입이 존재 하지 않음 --> rune(int32)로 문자 코드 값으로 표현
	//자주 사용하는 escape : \\ , \', \"

	var str1 string = "c:\\go"
	str2 := `c:\go_study`

	fmt.Println(str1)
	fmt.Println(str2)

	//바이트 수 : len 으로 구함
	// 실제 길이: utf8.RuneCountInString(str3) or len([]rune(str3)

}
