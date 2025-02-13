package main

import "fmt"

// "myFirstGoProject/pacote"

//bool:

// Inteiros :

// int int8 int16

//uint uint8 uint16 uintptr

// Type Alias:

// byte = uint8

//rune = int32

//float:

//float32 float64

//complexos:

//complex64 complex128

func main() {
	const x = 10

	TakeInt32(x)
	TakeInt64(x)
	TakeString("foo")
	TakeFloat(3.14)
	// x := 10084
	// s := strconv.FormatInt(int64(x), 10)

	// fmt.Println(s)
}

func TakeFloat(f float64) {
	fmt.Println(f)
}

func TakeString(s string) {
	fmt.Println(s)
}

func TakeInt32(x int32) {
	fmt.Println(x)
}

func TakeInt64(x int64) {
	fmt.Println(x)
}
