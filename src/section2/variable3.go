package main

import "fmt"

func main(){
	//짧은 선언
	//함수 안에서만 사용 (전역 X), 선언 후 할당 예외 발생
	shortVar1 := 3

	fmt.Println(shortVar1)

	//Example
	if i:= 10 ; i< 11{
		fmt.Println(i)
		i=11
	}
	shortVar1 = 4
	fmt.Println(shortVar1)

}