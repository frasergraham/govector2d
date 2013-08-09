package govector

import (
	"fmt"
	"math"
	"testing"
)

func TestSubMag(t *testing.T) {
	p1 := Point2d{10, 10}
	p2 := Point2d{7, 6}

	v := p1.Sub(p2)

	if v != (Vector2d{3, 4}) {
		t.Errorf("Sub Error")
	}

	m := v.Mag()

	if m != 5 {
		t.Errorf("Mag Error")
	}
}

func TestScale(t *testing.T) {
	v := Vector2d{3, 4}
	m := v.Mag()

	if m != 5 {
		t.Errorf("Mag Error")
	}

	m2 := v.Scale(2)

	fmt.Printf("%v\n", m2)

	if m2.Mag() != 10 {
		t.Errorf("Mag Error")
	}
}

func TestAdd(t *testing.T) {
	p := Point2d{10, 10}
	v := Vector2d{3, 4}
	m := p.Add(v)

	fmt.Printf("%v\n", m)

	if m.X != 13 && m.Y != 14 {
		t.Errorf("Add Error")
	}
}

func TestNormalize(t *testing.T) {
	v := Vector2d{3, 4}
	vn := v.Normalize()

	if vn.Mag() != 1 {
		t.Errorf("Normalize Error")
	}

	v0 := Vector2d{0, 0}
	vn0 := v0.Normalize()

	if vn0.Mag() != 0 {
		t.Errorf("Normalize Error")
	}

}

func TestDot(t *testing.T) {
	v1 := Vector2d{0, 5}
	v2 := Vector2d{5, 0}
	d := v1.Dot(v2)

	if d != 0 {
		t.Errorf("Dot Error")
	}

	v1 = Vector2d{1, 5}
	v2 = Vector2d{5, 0}
	d = v1.Dot(v2)

	if d < 0 {
		t.Errorf("Dot Error")
	}

	v1 = Vector2d{-1, 5}
	v2 = Vector2d{5, 0}
	d = v1.Dot(v2)

	if d > 0 {
		t.Errorf("Dot Error")
	}

	v1 = Vector2d{5, 5}
	v2 = Vector2d{5, 0}
	v1 = v1.Normalize()
	v2 = v2.Normalize()
	d = v1.Dot(v2)
	fmt.Printf("%v\n", d)

	if math.Abs(d-math.Acos(math.Pi/4)) < 0.0001 {
		t.Errorf("Dot Error")
	}
}

func TestIntersection(t *testing.T) {
	p1 := Point2d{0, 5}
	p2 := Point2d{5, 0}

	v1 := Vector2d{10, 0}
	v2 := Vector2d{0, 10}

	i, p := Intersection(p1, p2, v1, v2)
	fmt.Printf("%v %v\n", i, p)
	if !i {
		t.Errorf("Intersection Error")
	}

    p1 = Point2d{20, 5}
    p2 = Point2d{5, 0}

    v1 = Vector2d{10, 0}
    v2 = Vector2d{0, 10}

    i, p = Intersection(p1, p2, v1, v2)
    fmt.Printf("%v %v\n", i, p)
    if i {
        t.Errorf("Intersection Error")
    }

}
