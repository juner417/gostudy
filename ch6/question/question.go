package question

import (
	"fmt"
)

type Point struct{ x, y int }

func main() {
	//arr := [3]string{"h","e","l"}
	//s := []Point{}
	var s []Point
	// 위 두라인의 차이를 모르겠다...
	s = append(s, Point{5, 6})
	s = append(s, Point{100, 200})

	for i, v := range s {
		fmt.Println(i, v)
	}

}
