package main

import "fmt"

func main_10() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3

	fmt.Println("map:", m)
	fmt.Println("map K1:", m["k1"])

	delete(m, "k2")
	fmt.Println("map:", m)

}
