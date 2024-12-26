package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main_13() {
	i := 1

	fmt.Println("initial:", i)

	zeroval(1)

	fmt.Println("zeoval", i)

	zeroptr(&i)

	fmt.Println("zeroptr", i)

}
