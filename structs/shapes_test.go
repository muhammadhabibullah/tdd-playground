package structs

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{
			name: "Triangle",
			shape: RightTriangle{
				Base:   3.0,
				Height: 4.0,
			},
			hasPerimeter: 12.0,
		},
		{
			name: "Rectangle",
			shape: Rectangle{
				Width:  10.0,
				Height: 10.0,
			},
			hasPerimeter: 40.0,
		},
		{
			name: "Circle",
			shape: Circle{
				Radius: 10.0,
			},
			hasPerimeter: 62.83185307179586,
		},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			name: "Triangle",
			shape: RightTriangle{
				Base:   3.0,
				Height: 4.0,
			},
			hasArea: 6.0,
		},
		{
			name: "Rectangle",
			shape: Rectangle{
				Width:  10.0,
				Height: 10.0,
			},
			hasArea: 100.0,
		},
		{
			name: "Circle",
			shape: Circle{
				Radius: 10.0,
			},
			hasArea: 314.1592653589793,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
