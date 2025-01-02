package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "andré"
	u.idade = 25
	fmt.Println(u)

	u2 := usuario{
		"andré",
		25,
		endereco{
			logradouro: "exemplo",
			numero:     123,
		},
	}
	fmt.Println(u2)

	u3 := usuario{
		nome: "andré",
	}
	fmt.Println(u3)
}
