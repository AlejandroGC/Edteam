// package main es el paquete principal que me permite ejecutar mi aplicacion de EDteam
package main

// biblioteca que nos permite tener las bases y los instrumentos principales de Go
import "fmt"

func main() {
	// VARIABLES
	dog, cat := "perro", "gato"
	cat, face := "gatito", "cara"
	fmt.Println(dog, cat, face)


	// CONSTANTES
	const pi = 3.14
	fmt.Println(pi)

	// COMENTARIOS
	// /* comentario multilinea */
	// Posible documentar facilmente con terminal "go doc --all"


	// bool, string, numeric
	// var a bool = true
	// var b string = "Edteam"
	var c uint = 2000
	var d byte = 150
	// var e rune = -200
	// var f float64 = 56.55
	
	// Variable blank
	_ = 234
	var _ string = "test"

	suma := c + uint(d)

	// Verbos: %T , %v, %q
	fmt.Printf("Tipo: %T, Valor: %v \n", suma, suma)
	fmt.Printf("Tipo: %T, Valor: %v \n", c,c)


	// Operadores aritmeticas (), %, +, -, *, /
	var a = 4 + 2*3 
	fmt.Println(a)

	// Operadores de asigancion: =, +=, -=, *=, /=, %=
	var b = 10
	b += 2
	fmt.Println(b)

	//Operadores de Comparacion: >, <, >=, <=, !=, ==
	fmt.Println(4!= 4)
	// declaracion postincremento y post decremento ++, --


	var z = 3
	z--
	fmt.Println(z)

	// Operadores Logicos &&, ||
	var age = 30
	fmt.Println("edad:",  age >= 18 && age <= 60)
	fmt.Println("nino o viejo:", age < 18 || age > 60)
	// Unario: !
}