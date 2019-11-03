package geometry

import (
	"math"
)

//Point struct for method
type Point struct{ X, Y float64 }

//Distance legacy
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Distance method for Point type
// p is reciever
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//func main() {
//	p := Point{1, 2}
//	q := Point{4, 6}

//	fmt.Println(Distance(p, q)) //"5" call Distance func
//	fmt.Println(p.Distance(q))  //"5" call Distance method for Point struct
//}

//Path is path to connect between points
type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			// point type Distance를 호출한다.
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
