package structsmethodsinterfaces

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	perimeter = 2 * (rectangle.Height + rectangle.Width)
	return perimeter
}

func Area(rectangle Rectangle) (area float64) {
	area = rectangle.Height * rectangle.Width
	return area
}
