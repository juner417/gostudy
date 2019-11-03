package main

import (
	"fmt"
)

func main() {
	// array
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}

	// slice
	Q2 := months[4:7]
	summer := months[6:9]

	// 2quater
	fmt.Println("2nd quater---")
	for _, m := range Q2 {
		fmt.Println(m)
	}

	// summer
	fmt.Println("in summer--")
	for _, v := range summer {
		fmt.Println(v)
	}

	fmt.Printf("slice Q2: len-%d < cap-%d \n", len(Q2), cap(Q2))
	fmt.Printf("slice summer: len-%d < cap-%d \n", len(summer), cap(summer))

	for _, m := range Q2 {
		for _, v := range summer {
			if m == v {
				fmt.Printf("%s appears in both \n", m)
			}
		}
	}

	fmt.Printf("slice summer: %s \n", summer[:5])

	a := [...]int{0, 1, 2, 3, 4, 5} // 배열선언]
	//reverse(a) // type이 다르기 때문에 안됨 a 는 [6]int reverse 인자 s는 []int
	reverse(a[:])

	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("type for s: %T", s)
	reverse(s[:2])
	fmt.Printf("%d\n", s)
	reverse(s[2:])
	reverse(s)
	fmt.Printf("%d\n", s)

}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
