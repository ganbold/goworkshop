package main

import (
	"fmt"
	"reflect"
)

func main_9() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("apd:", s)

	t := []string{"g", "h", "i"}
	fmt.Println(reflect.TypeOf(t))
	t = append(t, "l")
	fmt.Println(t)
}
