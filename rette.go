package gogeom

import "fmt"
import "math"

type Line struct {
	m, q float64
}

func NewLine(m,q float64) Line {
	return Line{m,q}
}

//---

func (r Line) String() string {
	return fmt.Sprintf("y=%.3f * x + %.3f", r.m, r.q)
}

func (p1 Point) RettaPassante(p2 Point) Line {
	m := (p2.y - p1.y) / (p2.x - p1.x)
	q := p1.y - m * p1.x
	return NewLine(m, q)
}

func (r Line) PassaPer(p Point) bool {
	return p.y == r.m * p.x + r.q
}

func (r Line) DistOrigine() float64 {
	return math.Abs(r.q) / math.Sqrt(r.m*r.m + 1)
}


