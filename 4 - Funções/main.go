package main

import "fmt"

func main() {
	soma := somar(82, somar(5, somar(5, 91)))
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}
	f("texto")

	resultadosoma, resultadosub := calculosMatematicos(10, 15)
	fmt.Println(resultadosoma)
	fmt.Println(resultadosub)

	resultadosomentesoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadosomentesoma)
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}
