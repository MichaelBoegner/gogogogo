package structsmethodsinterfaces

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159265359 * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	perimeter = 2 * (rectangle.Height + rectangle.Width)
	return perimeter
}
