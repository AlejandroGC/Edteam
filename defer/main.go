package main

import (
	"fmt"
	"os"
)

func main(){
	// a := 5
	// defer fmt.Println("Defer:", a)

	// a = 10
	// fmt.Println(a)
	
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo: %v", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("Hello Edteam"))
	if err != nil {
		fmt.Println("Error al escribir el archivo: %v", err)
		return
	}

}