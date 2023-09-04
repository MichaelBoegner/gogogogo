package structsmethodsinterfaces

func Perimeter(height, width int) (perim int) {
	perim = 2 * (height + width)
	return perim
}

func Area(height, width float64) (area float64) {
	area = height * width
	return area
}
