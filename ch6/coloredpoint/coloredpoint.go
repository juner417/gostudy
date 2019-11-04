package coloredpoint

import (
	"image/color"
	"math"
)

//Point ...
type Point struct{ X, Y float64 }

//ColoredPoint ...
type ColoredPoint struct {
	Point //anonymous field
	Color color.RGBA
}

//PColoredPoint ...
type PColoredPoint struct {
	*Point //anonymous field, 이것도 Point struct 필드와 메소드가 승격됨 refered(간접적으로)
	Color  color.RGBA
}

//DColoredPoint ...
type DColoredPoint struct {
	Point
	color.RGBA
	//
}

//Distance ...
func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

//ScaleBy ...
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
