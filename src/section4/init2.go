package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init1 Method Start!")
}

func init() {
	fmt.Println("Init2 Method Start!")
}
func init() {
	fmt.Println("Init3 Method Start!")
}
func init() {
	fmt.Println("Init4 Method Start!")
}
func init() {
	fmt.Println("Init5 Method Start!")
}

func main() {
	//가장 나중에 import된 파일의 init부터 차례대로 호출된다 (!!!! import되는 순간 import되는 파일의 init 함수가 실행된다!!! )
	//init 함수 여러개 가능 --> 중요하지는 않음

}
