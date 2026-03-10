package main

import "fmt"

func main() {

	var c int

	fmt.Print("Input Celsius: ")
	fmt.Scanln(&c)

	f := (c * 9 / 5) + 32
	r := c * 4 / 5
	k := c + 273

	fmt.Println("Derajat Reamur:", r)
	fmt.Println("Derajat Fahrenheit:", f)
	fmt.Println("Derajat Kelvin:", k)

}
