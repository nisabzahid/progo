package main

import "fmt"

type Rectangle struct {
	Height int
	Width  int
}

func (r Rectangle) Area() int {
	return r.Height * r.Width
}

type Box struct {
	Rectangle
	depth int
}

func (b Box) Volume() int {
	return b.Area() * b.depth
}

func main() {

	b := Box{Rectangle{5, 5}, 5}
	fmt.Printf("%+v\n", b)
	fmt.Println("voluem ", b.Volume())
	fmt.Printf("Hello, %s!\n")

}
