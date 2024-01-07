package shapes

import (
	"testing"
)

func checkArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("shape.Area() of %#v = \"%f\"; want \"%f\"", shape, got, want)
	}
}

func checkPerimeter(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Perimeter()
	t.Log(got)
	if got != want {
		t.Errorf("shape.Area() of %#v = \"%f\"; want \"%f\"", shape, got, want)
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 40},
		{Circle{10.0}, 62.83185307179586},
		{Triangle{Base: 10.0, Height: 6.0}, 36},
	}

	for _, tt := range perimeterTests {
		checkPerimeter(t, tt.shape, tt.want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Width: 10.0, Height: 10.0}, 100},
		{Circle{Radius: 10.0}, 314.1592653589793},
		{Triangle{Base: 12.0, Height: 6.0}, 36},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}
