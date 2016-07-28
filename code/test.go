package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

/*
4 : 0
6 : 0
5 : 0
7 : 0
0 : 0
8 : 0
2 : 0
3 : 0
1 : 0
9 : 0
3 : 1
9 : 1
7 : 1
9 : 2
8 : 1
0 : 1
4 : 1
5 : 1
*/
