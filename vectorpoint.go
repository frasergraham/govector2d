package govector

import (
	"math"
)

type Vector2d struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Point2d struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

const epsilon = 1e-7

func (p1 Point2d) sub(p2 Point2d) Vector2d {
	v := Vector2d{p1.X - p2.X, p1.Y - p2.Y}
	return v
}

func (v Vector2d) mag() float64 {
	return math.Abs(math.Sqrt(v.X*v.X + v.Y*v.Y))
}
