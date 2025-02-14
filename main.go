package main

import (
	"fmt"
	"math"
)

func main() {
	if x := math.Sqrt(4); x < 1 {
		fmt.Println(x)
	} else if x < 1 {
		fmt.Println("Maior que zero")
	} else {
		fmt.Print("caiu aqui")
	}
}
