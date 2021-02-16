// switch문 (1)
package main

import "fmt"

func main(){
	// 제어문 (조건문) switch
	//Switch 뒤 표현식은 생략이 가능
	// case 뒤 표현식은 사용 가능
	// 자동 break 때문에 fallthrouth 가 존재
	// Tpye 분기 --> 값이 아닌 변수 type 으로 분기 가능

	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a ,"는 양수")
	}

	switch b := 26; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b ,"는 양수")
	}

	switch c := "go"; c {
	case "go":
		fmt.Println("Go")
	case "java":
		fmt.Println("java")
	default:
		fmt.Println("no")
	}

	switch c := "go"; c+"lang"{
	case "golang":
		fmt.Println("Golang")
	case "java":
		fmt.Println("java")
	default:
		fmt.Println("no")
	}

	switch i,j := 20,30; {
	case i < j:
		fmt.Println("i<j")
	case i == j:
		fmt.Println("i==j")
	}

}
