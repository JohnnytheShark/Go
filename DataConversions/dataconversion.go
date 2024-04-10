package main

import "fmt"

// func main() {
// 	var i int = 20
// 	var f float64 = float64(i)
// 	fmt.Println(i)
// 	fmt.Println(f)
// }

func main(){
	const value = 20
	var i int = value
	var f float64 = float64(value)

	fmt.Println("Integer i =", i)
	fmt.Println("Float f =",f)
}
