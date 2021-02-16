package main

import "fmt"

func main() {
	//Go: 포인터 지원(c)
	//변수의 지역성, 연속된 메모리 참조성 , 힙 스
	//포인터 주소값은 임의 변경 불가능
	//* 에스터리스크 로 사용
	// nil로 초기화된다

	var a *int            //방법 1
	var b *int = new(int) // 방법2 정석!!!!!!
	fmt.Println(a)
	fmt.Println(b)

	i := 710
	fmt.Println(&i) // &i 7이 저장된 메모리의 주소

	fmt.Println()
	a = &i
	b = &i
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*a) // 역 참

	fmt.Println()

	*a = 1000

	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(*b)

}
