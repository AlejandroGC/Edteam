package main 

import (
	"fmt"
	//"io/ioutil"
	"errors"
)

func main(){
// 	content, err := ioutil.ReadFile("things.txt")
// 	if err != nil {
// 		fmt.Printf("Ocurrio un error: %v", err)
// 		return
// 	}
//  fmt.Println(string(content))
result, err := division(10,2)
if err != nil {
	fmt.Printf("Ocurrio un error: %v", err)
	return
}
fmt.Println(result)

}

func division(dividendo, divisor int)(result int, err error) {
	if divisor == 0 {
		err = errors.New("No puedes dividir por 0") 
		return
	}
	result = dividendo / divisor
	return
}