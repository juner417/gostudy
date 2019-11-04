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

//Distance is method for Path type
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

//-- 6.1 end
//-- 6.2 start

//ScaleBy 주어진 값만큼 좌표 수정
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

//LocalScaleBy not point reciever
func (p Point) LocalScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

//func main() {
//	p := Point{1, 2}
//	fmt.Println(p) //{1,2}
//	p.ScaleBy(3)
//	fmt.Println(p) //{3,6}
//	p.LocalScaleBy(2)
//	fmt.Println(p) //{3,6}
//}
//-- 6.2 end
//-- 6.4 start
//Add ...
func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }

//Sub ...
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

//TranslateBy ...
func (path Path) TranslateBy(offset Point, add bool) Path {
	var op func(p, q Point) Point // op 변수에 익명 함수 선언
	if add {
		op = Point.Add // 메소드 표현식
	} else {
		op = Point.Sub // 메소드 표현식
	}
	for i := range path {
		// path[i].Add(offset) 또는 path[i].Sub(offset)을 호출한다.
		path[i] = op(path[i], offset)
	}
	// fmt.Printf("%#v, %v \n", op, path) //for debuging
	return path
}
