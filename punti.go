package gogeom

import "fmt"
import "math"

type Point struct {
	x, y float64
}

func NewPoint(x,y float64) Point {
	return Point{x,y}
}

//----

func (p Point) String() string {
	return fmt.Sprintf("(%.3f, %.3f)", p.x, p.y)
}

func (p Point) DistOrigine() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p1 Point) Dist(p2 Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}


func (p Point) GetX() float64 {
	return p.x
}

func (p Point) GetY() float64 {
	return p.y
}



