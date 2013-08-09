package govector

import (
    "math"
)

type vector2d struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type point2d struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

const epsilon = 1e-7

func (p1 point2d) sub(p2 point2d) vector2d {
	v := vector2d{p1.X - p2.X, p1.Y - p2.Y}
	return v
}

func (v vector2d) mag() float64 {
	return math.Abs(math.Sqrt(v.X*v.X + v.Y*v.Y))
}
