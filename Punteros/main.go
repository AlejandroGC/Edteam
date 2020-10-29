package main

import "fmt"

func main(){
	fruit := "manzana"
	var p *string
	p = &fruit 
	*p = "piña"
	fmt.Printf("Tipo: %T, Valor: %s, Direccion: %v \n", fruit, fruit, &fruit)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciacion: %s \n", p, p, *p)
}