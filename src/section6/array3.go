package main

import "fmt"

func main() {

	//값 복사
	arr1 := [5]int{1, 10, 100, 1000, 10000}
	arr2 := arr1

	fmt.Println(arr2)

	fmt.Printf("%p %v \n", &arr1, arr1)
	fmt.Printf("%p %v \n", &arr2, arr2)
	//arr1 , arr2 주소값이 다르다는것을 알 수 있다 == 값을 복사 했다.

}
