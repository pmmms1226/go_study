package main

import "fmt"

func main() {

	cnt := increseCnt()
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())

}

func increseCnt() func() int {
	n := 0 //지역변수 캡처 --> 소명할지 않는다.
	return func() int {
		n += 1
		return n
	}
}
