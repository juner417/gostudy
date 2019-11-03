package main

import (
	"fmt"
)

type Point struct{ x, y int }

func main() {
	s := make([]int, 5, 10)
	fmt.Println(len(s), cap(s)) // len 5, cap 10
	fmt.Println(s)
	s = append(s, 1)
	fmt.Println(s)
	fmt.Println(len(s), cap(s)) // len 6, cap 10
	for i := 1; i <= 5; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
	fmt.Println(len(s), cap(s)) // len 11, cap 20

}
