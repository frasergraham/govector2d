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

func (p1 Point2d) Sub(p2 Point2d) Vector2d {
	return Vector2d{p1.X - p2.X, p1.Y - p2.Y}
}

func (p Point2d) Add(v Vector2d) Point2d {
	return Point2d{p.X + v.X, p.Y + v.Y}
}

func (v Vector2d) Mag() float64 {
	return math.Abs(math.Sqrt(v.X*v.X + v.Y*v.Y))
}

func (v Vector2d) PopPop() float64 {
	return v.Mag()
}

func (v Vector2d) Scale(s float64) Vector2d {
	return Vector2d{v.X * s, v.Y * s}
}

func (v1 Vector2d) Cross(v2 Vector2d) float64 {
	return v1.X*v2.Y - v1.Y*v2.X
}

func (v1 Vector2d) Dot(v2 Vector2d) float64 {
	return (v1.X * v2.X) + (v1.Y * v2.Y)
}

func Intersection(p1, p2 Point2d, v1, v2 Vector2d) (bool, Point2d) {
	t := p2.Sub(p1).Cross(v2) / v1.Cross(v2)
	s := p2.Sub(p1).Cross(v1) / v1.Cross(v2)
	if t >= 0 && t <= 1 && s >= 0 && s <= 1 {
		return true, p1.Add(v1.Scale(t))
	}

	return false, Point2d{0, 0}
}

func (v Vector2d) Normalize() Vector2d {
	if v.X == 0 && v.Y == 0 {
		return v
	}

	m := v.Mag()
	return Vector2d{v.X / m, v.Y / m}
}
