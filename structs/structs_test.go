package structs

import (
	"hello/testhelper"
	"math"
	"testing"
)

type Rectangle struct {
	W, H float32
}

type Circle struct {
	Radius float32
}

type Triangle struct {
	W, H float32
}

type Shape interface {
	Area() float32
}

func Perimeter(r Rectangle) float32 {
	return 2 * (r.W + r.H)
}

func (r Rectangle) Area() float32 {
	return r.W * r.H
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float32 {
	return t.W * t.H / 2
}

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := Perimeter(r)
	var wanted float32 = 40.0
	testhelper.AssertFloat(t, wanted, got)
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, s Shape, wanted float32) {
		t.Helper()
		got := s.Area()
		testhelper.AssertFloat(t, wanted, got)
	}

	shapes := []struct {
		Shape   Shape
		HasArea float32
	}{
		{Shape: Rectangle{10.0, 10.0}, HasArea: 100.0},
		{Shape: Circle{10.0}, HasArea: 314.16},
		{Shape: Triangle{10.0, 5.0}, HasArea: 25.0},
	}

	for _, s := range shapes {
		checkArea(t, s.Shape, s.HasArea)
	}

}
