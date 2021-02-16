package main

import "fmt"

func main(){
	//제어문(조건문)
	//if문은 반드시 Boolean 으로 검사해야한다. --> go 에서는 1,0 사용 불가 무조건 True, False
	// 소괄호를 사용하지 않는다.
	var a int = 20
	b := 20
	if a >= 15 {
		fmt.Println("15이상")
	}

	if b>= 25{
		fmt.Println("25")
	}

	if c := true; c{
		fmt.Println("hello world")
	}
}
