package main

import "fmt"

type Temperature struct {
	Celcius float64
}

func (t Temperature) ToFahrenheit() float64 {
	return 9./5*t.Celcius + 32
}

func (t Temperature) ToKelvin() float64 {
	return t.Celcius + 273.0
}

func main() {
	c := Temperature{28}

	fmt.Println("Celcius:", c.Celcius)
	fmt.Println("Kelvin:", c.ToKelvin())
	fmt.Println("Fahrenheit:", c.ToFahrenheit())
}
