package main

import "fmt"

type Rectangle struct {
	width, height int
}

func (r *Rectangle) area() int {
	return r.height * r.width
}

func (r *Rectangle) perimeter() int {
	return 2*r.width + 2*r.height
}

func main_16() {
	r := Rectangle{width: 10, height: 5}

	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perimeter())

	rp := &r

	fmt.Println("Sides:", rp.width, " x ", rp.height)
	fmt.Println("Area:", rp.area())
	fmt.Println("Perimeter:", rp.perimeter())

}
