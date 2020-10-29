package main

import "fmt"

func main(){
	emoji := "perro"
	switch emoji {
	case "perro", "gato":
		fmt.Println("es un animal")
	case "Alex", "Daniel":
		fmt.Println("es una persona")
	default:
		fmt.Println("es un alien")
	}
}