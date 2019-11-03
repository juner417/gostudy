package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, playground")

	var a [3]int = [3]int{4,5,6}

	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d \n", v)
	}

	var q [3]int = [3]int{1,2,3}
	
	for _, v := range q{
		fmt.Printf("%d \n", v)
	}
	
	//var ss [10]string = [10]string{"kubernetes"}
	var ss string = "kubernetes"
	
	
	for i, v := range ss {
		fmt.Printf("%d %s \n", i, strconv.Itoa(int(v)))
	}
	
	p := [...]int{1,2,3}
	
	for _, v := range p {
		fmt.Printf("%d \n", v)
	}
	
	w := [2]int{1,2}
	x := [...]int{1,2}
	y := [2]int{1,3}
	fmt.Println(w == x, w == y, x == y)
	//z := [3]int{1,2}
	//fmt.Println(w == z)
}

