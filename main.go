package main

import (
	"fmt"
)

func slices() {
	var s []string
	fmt.Println("uninit:", s, "s == nil", s == nil, "len(s) == 0", len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "apple"
	s[1] = "banana"
	s[2] = "carrot"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	s = append(s, "Rick")
	s = append(s, "Morty", "Summer")
	fmt.Println("appended:", s)

}

func main() {
	slices()
}
