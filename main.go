package main

import "fmt"

func main() {
	doDefer()
}

func doDefer() {
	x := 10
	defer func(y int) {
		fmt.Println(y)
	}(x)

	x = 50
	fmt.Println(x)
}

// func doDefer() {
// 	defer fmt.Println(3)
// 	defer fmt.Println(2)
// 	fmt.Println(1)
// }

// func doDefer() int {
// 	defer fmt.Println("World")
// 	fmt.Println("Hello")
// 	return 10
// }
