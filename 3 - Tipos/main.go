package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	var numero2 uint = 20
	fmt.Println(numero2)

	var numeroreal1 float64 = 123.40
	fmt.Println(numeroreal1)

	char := 'A'
	fmt.Println(char)

	var textoinicial string
	fmt.Println(textoinicial)

	var numeroinicial int
	fmt.Println(numeroinicial)

	var booleano bool = true
	fmt.Println(booleano)

	var booleanoinicial bool
	fmt.Println(booleanoinicial)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}
