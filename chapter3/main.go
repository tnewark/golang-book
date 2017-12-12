package main

import "fmt"

var x = "Hello, World"
const c = "this is a constant"

func main() {
	var (
		a = 5
		b = 10
		c = 15
	)
	var d = a + b + c
	fmt.Println("d = " , d)
	fmt.Println(x)

	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)

	x = "hello"
	y := "world"
	fmt.Println(x == y)

	fmt.Print("enter a farenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input-32) * 5/9

	fmt.Println(output)
}