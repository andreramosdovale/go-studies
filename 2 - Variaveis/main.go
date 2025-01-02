package main

func main() {
	var variavel1 string = "variavel1"
	println(variavel1)

	variavel2 := "variavel2"
	println(variavel2)

	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)
	println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "variavel6"
	println(variavel5, variavel6)

	const constante1 string = "constante1"
	println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	println(variavel5, variavel6)
}
