package main

import "fmt"

func main() {

	// Variable Declaration With Initial Value
	var name string = "abderrahmane" // type is string
	fmt.Println(name)

	var x, y int = 213, 99 // type is integer
	fmt.Println(x, y)

	var even bool = true // type is boolean
	fmt.Println(even)

	var pi float32 = 3.14
	fmt.Println(pi)

	var names = [...]string{"abderrahmane", "nadir", "ayoub", "aymen"}
	fmt.Println(names)

	// Variable Declaration Without Initial Value
	var a int
	fmt.Println(a)

	var b string
	fmt.Println(b)

	var c bool
	fmt.Println(c)

	var products [5]int
	fmt.Println(products)

	// Value Assignment After Declaration
	var number int
	number = 213
	fmt.Println(number)

	// Variable declaration and initialization
	newNumber := 286
	fmt.Println(newNumber)

}
