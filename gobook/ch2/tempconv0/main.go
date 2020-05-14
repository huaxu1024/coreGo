package main

import "fmt"

func main() {

	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF - CToF(FreezingC))

	var c Celsius
	var f Fahrenheit

	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	fmt.Println(c == Celsius(f))
	fmt.Println(c.String())
}

func init() {
	fmt.Println("-----init------")
}
