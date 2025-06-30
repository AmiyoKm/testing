package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Heigh float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Heigh)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Heigh
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return roundTo2(2 * math.Pi * c.Radius)
}

func (c Circle) Area() float64 {
	return roundTo2(math.Pi * c.Radius * c.Radius)
}

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) Area() float64 {
	return (roundTo2(t.Base * t.Height)) / 2
}

func roundTo2(f float64) float64 {
	return math.Round(f*100) / 100
}
