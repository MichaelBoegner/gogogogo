package main

import (
    "fmt"
    "math/rand"
	"time"
)

func main() {
	var ageJohn, agePaul int
	
	rand.Seed(time.Now().UTC().UnixNano())
	ageJohn, agePaul = rand.Intn(110), rand.Intn(110)
	
	fmt.Println("ages for John and Paul respectively:", ageJohn, agePaul)
	
	if ageJohn > agePaul {
		fmt.Println("It is true that John is older than Paul")
	} else if ageJohn == agePaul {
		fmt.Println("John and Paul are the same age")
	} else {
		fmt.Println("It is true that Paul is older than John")
	}
}