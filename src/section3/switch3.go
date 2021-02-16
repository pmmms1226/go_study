package main

import "fmt"

func main(){
	a := 30 / 15
	switch a{
	case 2,4,6:
		fmt.Println(a) // i가 2,4,6 인 경우
	case 1,3,5:
		fmt.Println(a)
	}
 	//fallthrough 는 다음 case로 넘겨라
	switch e := "go" ; e{
	case "java":
		fmt.Println("java!")
		fallthrough
	case "go":
		fmt.Println("go!")
		fallthrough
	case "python":
		fmt.Println("python!")
		fallthrough
	case "ruby":
		fmt.Println("ruby!")
	case "c":
		fmt.Println("c!")
	}
}