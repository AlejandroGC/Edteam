package main

import "fmt"

func main(){
	// set := [7]string {"leon", "caballo", "vaca", "mariposa", "pajaro", "avioneta", "avion"}
	// animals := set[:5]
	// fly := set[3:]
	// fmt.Println("Array:", animals)
	// fmt.Println("Voaldores:", fly)
	// fmt.Println("Voaldores:", set[:])

	// num := [5]int {4,5,6,7,8}
	// nums := num[1:3]
	// nums = append(nums,18,23)

	// fmt.Println("num", num)
	// fmt.Println("nums", nums)
	// fmt.Println("tamaño", len(nums))
	// fmt.Println("cantidad", cap(nums))

	// ---------------------------------------
	// fruits := make([]string, 0 ,3)

	// fruits = append(fruits, "piña","manzana","uva","naranja")

	// fmt.Println("fruits", fruits)
	// fmt.Println("tamaño", len(fruits))
	// fmt.Println("capacidad", cap(fruits))

	// -----------------------------------------

	var fruits [] string
	fruit := [] string {}
	fmt.Println("fruits", fruits)
	fmt.Println("tamaño", len(fruits))
	fmt.Println("capacidad", cap(fruits))
	fmt.Println(fruits == nil)
	fmt.Println(fruit == nil)

}