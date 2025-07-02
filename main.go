package main

import "fmt"

type Rsvp struct {
	name, email, phone string
	WillAttend         bool
}

var responses = make([]*Rsvp, 0, 10)

func main() {

	fmt.Println("TODO: Add some features")

}
