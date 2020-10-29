package main

import "fmt"
import "strings"

func main(){
	texto := "aLEjanDRo"
	minusc, mayusc := convert(texto)
	fmt.Println(minusc,mayusc)

}

func convert(text string) (string, string){
	min := strings.ToLower(text)
	may := strings.ToUpper(text)

	return min, may
}

// func sum(num1, num2 int) int{
// 	return num1 + num2
// }

// func change(value *string){
// 	*value = "emoji"
// }


// func hello(firstName string, lastName string){
// 	fmt.Printf("Hola %s %s", firstName, lastName)
// }