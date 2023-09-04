package structsmethodsinterfaces

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return 0
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 0
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	perimeter = 2 * (rectangle.Height + rectangle.Width)
	return perimeter
}

func Area(rectangle Rectangle) (area float64) {
	area = rectangle.Height * rectangle.Width
	return area
}
