package main

func main() {

	//길이가 가변, 참조 값 타입
	//길이가 가변형 이기때문에 초기화가 안된다.
	//배열처럼 생성 가능, make함수로 생성 make(자료형, 길이, 가변용량)

	//var slice1 []int
	//slice2 := []int{}
	//slice3 := []int{1, 2, 3, 4, 5}
	//
	////예제2 make는 초기화가 가능함 make를 더 많이 사용함
	//var slice5 []int = make([]int, 5, 10)
	//var slice6 = make([]int, 5) //제일 많이 사용함
	//slice7 := make([]int, 6, 100)

	var slice9 []int //nil 값이다 --> make로 초기화를 해주던 해야함
	slice9[1] = 1

}
