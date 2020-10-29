package main

import "fmt"

func main(){
	emoji := "cactus"
	if emoji == "cactus"{
		fmt.Println("es un cactus")
	} else if emoji == "carita"{
		fmt.Println("es una carita")
	} else {
		fmt.Printf("el emoji es %s", emoji)
	}

	
	if emoji := "gato"; emoji == "cactus" {
		fmt.Println("es un cactus")
	} else if emoji == "carita"{
		fmt.Println("es una carita")
	} else {
		fmt.Printf("el emoji es %s", emoji)
	}
}