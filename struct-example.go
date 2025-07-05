package main

import "fmt"

type Rectangle struct {
	Height int
	Width  int
}

// Enlarge demonstrates pass-by-value for the receiver.
// Changes made to r inside this method are made to a copy of the original Rectangle,
// so the original Rectangle instance is not modified.
func (r Rectangle) Enlarge(factor int) {
	r.Height = r.Height * factor
	r.Width = r.Width * factor
}

// EnlargeP demonstrates pass-by-pointer for the receiver.
// Changes made to r inside this method are made to the original Rectangle instance
// because r is a pointer to that instance.
func (r *Rectangle) EnlargeP(factor int) {
	r.Height = r.Height * factor
	r.Width = r.Width * factor
}

func main() {

	rect := Rectangle{Height: 10, Width: 10}

	fmt.Println("Initial rectangle:", rect) // Expected: {10 10}

	rect.Enlarge(2)
	// After calling Enlarge (value receiver), rect should be unchanged
	// because Enlarge operates on a copy.
	fmt.Println("After Enlarge(2):", rect) // Expected: {10 10}

	rect.EnlargeP(5)
	// After calling EnlargeP (pointer receiver), rect should be changed
	// because EnlargeP operates on the original instance via a pointer.
	fmt.Println("After EnlargeP(5):", rect) // Expected: {50 50}

}
