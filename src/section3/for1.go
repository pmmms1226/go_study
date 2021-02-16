package main

import "fmt"

func main(){
	//반복문 - for
	//go 에서 유일하게 제공되는 반복문
	// 다양한 사용법을 숙지 해야함

	for i := 0; i < 5; i++{
		fmt.Println(i)
	}
    ////무한루프
	//for {
	//
	//}

	loc := []string{"seoul", "busan", "Incheon"}
	for index, name := range loc{ //필요없는 값은 _ 이거로 생략가능
		fmt.Println("index: ", index, "name: ", name)
	}
}
