package main

import "fmt"

func main() {
	sum := 1 + 2
	sub := 1 - 2
	div := 1 / 2
	mult := 1 * 2
	mod := 1 % 2
	fmt.Println(sum, sub, div, mult, mod)

	fmt.Println("----------")

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	fmt.Println("----------")

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	fmt.Println("----------")

	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)
	numero--
	numero -= 20
	fmt.Println(numero)
	numero *= 3
	fmt.Println(numero)
	numero /= 2
	fmt.Println(numero)

	fmt.Println("----------")
}
