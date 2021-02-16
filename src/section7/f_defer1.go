package main

import "fmt"

func ex_f1() {
	fmt.Println("f1 : start1")
	defer ex_f2() //마지막 호출
	fmt.Println("f1 : end")
}
func ex_f2() {
	fmt.Println("f2 : called!")
}

func main() {
	//Defer 함수 실행(지연, 지연했다가 나중에 실행)
	//Defer를 호출한 함수가 종료되기 직전에 호출 된다.
	//타 언어의 finally 문과 비슷
	// 주로 리소스 반환(DB 연결 종료), 파일 닫기
	ex_f1()

}
