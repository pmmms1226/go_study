package main

import (
	"fmt"
	_ "section4/lib" //나중에 사용할 수 도 있으니까 남겨두는 방식
	checkUp "section4/lib2"
)

func main() {
	fmt.Println("10보다 큰 수 ", checkUp.CheckNum1(1))

}
