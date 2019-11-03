package main

import (
	"fmt"

	"github.com/juner417/gostudy/ch6/geometry"
)

func main() {

	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}

	//"5" call Distance func
	//패키지 수준 함수 geometry.Distance 를 선언
	fmt.Println(geometry.Distance(p, q))
	fmt.Printf("type: %T\n", geometry.Distance(p, q))
	//"5" call Distance method for Point struct
	//Point type의 메소드를 선언 Point.Distanc
	fmt.Println(p.Distance(q))
	fmt.Printf("type: %T\n", q.Distance(q))
}
