package main

import "fmt"

func main() {

	//슬라이스 추가 및 병합 --> 용량은 적절하게 잡아주는게 좋다.

	s1 := []int{1, 2, 3, 4, 5}

	s1 = append(s1, 6, 7, 8, 9, 10)

	s2 := []int{6, 1, 2, 3, 4, 5}

	s2 = append(s1, s2...) // 슬라이스 어팬드는 ... 을 붙여야한다.

	fmt.Println(s1)

	fmt.Println(s2)

	//

	s4 := make([]int, 0, 5)

	//용량은 초과하면 2배씩 증가한다
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
	}
	fmt.Println(s4)
}
