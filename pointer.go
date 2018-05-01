package main 

import "fmt"

func swap(num1 *int, num2 *int) {
	util := *num1 
	*num1 = *num2 
	*num2 = util
}

func main() {
	a := 505
	b := 0

	fmt.Println(a, b)

	swap(&a, &b)

	fmt.Println(a, b)
}
