package main

import (
	"fmt"
	"strings"
)

func main() {
	//예제 1 (결합: 일반연산)

	// + 연산으로 string 결합할 수 있음
	/// 예제2 (결합: Join)

	strSet := []string{} //string 배열선언
	strSet = append(strSet, "hello world")
	strSet = append(strSet, "hello2 world2")

	fmt.Println("ex2: ", strings.Join(strSet, "-----"))

}
