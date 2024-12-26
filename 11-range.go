package main

import "fmt"

func main_11() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for index, num := range nums {
		fmt.Println(index)
		sum += num
	}

	fmt.Println("sum:", sum)

	kvs := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
	}

	for k, v := range kvs {
		fmt.Println("key:", k, ", value:", v)
	}

}
