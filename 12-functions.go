package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func multi_return() (int, int) {
	return 2, 3
}

func variadic_sum(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func main_12() {
	res := plus(1, 2)
	fmt.Println("sum:", res)

	a, b := multi_return()
	res = plus(a, b)

	fmt.Println("multi sum:", res)

	res = variadic_sum(1, 2, 3, 4, 5)

	fmt.Println("variadic sum:", res)
}
