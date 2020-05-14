package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZerC Celsius = -272.16
	FreezingC Celsius	= 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5+32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32)*5/9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g度", c)
}