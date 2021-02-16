package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	rand.Seed(time.Now().Unix())
	switch i := rand.Intn(100);{
	case i >= 50 && i < 100:
		fmt.Println(i)
	}
}
