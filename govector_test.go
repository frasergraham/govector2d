package govector

import (
	// "fmt"
	"testing"
)

func TestSub(t *testing.T) {
	p1 := point2d{10, 10}
	p2 := point2d{5, 5}

	v := p1.sub(p2)

	if v != (vector2d{5, 5}) {
		t.Errorf("Sub Error")
	}
}

func TestMag(t *testing.T) {
	p1 := point2d{10, 10}
	p2 := point2d{7, 6}

	v := p1.sub(p2)
	m := v.mag()

	if m != 5 {
		t.Errorf("Mag Error")
	}
}
