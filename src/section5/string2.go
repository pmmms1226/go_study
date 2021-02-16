package main

import "fmt"

func main() {
	//문자열: UTF-8 인코딩 (유니코드 문자 집합) -> 바이트 수 주의!

	//에제1
	var str1 string = "Golang"
	//var str2 string = "World"
	var str3 string = "고프로그래밍"

	fmt.Println("ex1: ", str1[0], str1[1])

	conStr := []rune(str3) //한글이 제대로 나온다
	fmt.Printf(" %c %c %c %c %c %c \n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5])

	for i, char := range str1 {
		fmt.Printf("ex3: %c(%d)\t", char, i)
	}

}
