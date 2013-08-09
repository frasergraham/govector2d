package govector

import (
	"fmt"
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
}
