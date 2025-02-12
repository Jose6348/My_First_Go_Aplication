package main

import (
	"fmt"
	"myFirstGoProject/pacote"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(pacote.Bar)
}

//Nomes Publicos ou Privados

//letra minúscula privado
//letra Maiúscula pública

//Nomes

var Foo string

const Bar string = "bar"

type MeuTipo struct{}
