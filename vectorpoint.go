package govector

import (
	"math"
)

type Rect2d struct {
	X1 float64 `json:"x1"`
	Y1 float64 `json:"y1"`
	X2 float64 `json:"x2"`
	Y2 float64 `json:"y2"`
}

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
	return Vector2d{p1.X - p2.X, p1.Y - p2.Y}
}

func (p Point2d) add(v Vector2d) Point2d {
	return Point2d{p.X + v.X, p.Y + v.Y}
}

func (v Vector2d) mag() float64 {
	return math.Abs(math.Sqrt(v.X*v.X + v.Y*v.Y))
}

func (v Vector2d) scale(s float64) Vector2d {
	return Vector2d{v.X * s, v.Y * s}
}
