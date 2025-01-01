package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	auxiliar.Escrever()
	auxiliar.Escrever2()
	fmt.Println("escrevendo do arquivo main")
	err := checkmail.ValidateFormat("andre.vale@a.com")
	fmt.Println(err)
}
