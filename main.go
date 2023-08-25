package main

func main() {
	x := 2
	y := 4
	Add(x, y)
	Subtract(x, y)
}

func Add(x, y int) (res int) {
	return x + y
}

func Subtract(x, y int) (res int) {
	return x - y
}
