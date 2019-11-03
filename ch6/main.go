package main

import (
	"fmt"

	"github.com/juner417/gostudy/ch6/geometry"
)

func main() {

	//g := &geometry.Point{1, 2}
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}

	//fmt.Printf("_ %T \t %#v \n", p, p)
	//"5" call Distance func
	//패키지 수준 함수 geometry.Distance func를 선언(p,q를 인자로 받음)
	fmt.Println(geometry.Distance(p, q))
	fmt.Printf("type: %T\n %#v\n", geometry.Distance, geometry.Distance)
	//"5" call Distance method for Point struct
	//Point type의 메소드를 선언 Point.Distanc(p객체의 slector로 호출)
	fmt.Println(p.Distance(q))
	fmt.Printf("type: %T\n %#v\n", q.Distance, q.Distance)

	perim := geometry.Path{
		geometry.Point{1, 1},
		geometry.Point{5, 1},
		geometry.Point{5, 4},
		geometry.Point{1, 1},
	}
	fmt.Println(perim.Distance()) // 12
}
