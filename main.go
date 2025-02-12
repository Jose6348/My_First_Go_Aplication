package main

import "fmt"

// "myFirstGoProject/pacote"

func main() {
	// fmt.Println("Hello world")

	// fmt.Println(pacote.Bar)

	// fmt.Println(somar(1, 2))
	// a, b := swap(10, 20)
	// fmt.Println(a, b)

	// res, rem := dividir(10, 3)
	// fmt.Println(res, rem)

	// x := somar(2)(1)
	// fmt.Println(x)

	fmt.Println(somar(10, 10, 10))
}

func somar(nums ...int) int {
	var out int
	for _, n := range nums {
		out += n
	}
	return out
}

// func somar(a int) func(int) int {
// 	return func(b int) int {
// 		return a + b
// 	}
// }

// chmando funções

// func somar(a, b int) int {
// 	return a + b
// }

// func swap(a, b int) (int, int) {
// 	return b, a
// }

// func dividir(a, b int) (res int, rem int) {
// 	res = a / b
// 	rem = a & b
// 	return res, rem
// }

//Nomes Publicos ou Privados

//letra minúscula privado
//letra Maiúscula pública

//Nomes

var Foo string

const Bar string = "bar"

type MeuTipo struct{}
