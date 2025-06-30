package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %0.2f want %0.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Perimeter()
		want := 62.83

		if got != want {
			t.Errorf("got %0.2f want %0.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.16},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %0.2f want %0.2f", tt.shape, got, tt.want)
			}
		})
	}
}
