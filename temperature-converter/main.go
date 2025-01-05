package main

import "fmt"
import "temperature-converter/tempconv"

func main() {
	var celsius tempconv.Celsius
	var fahrenheit tempconv.Fahrenheit
	celsius = 40
	fahrenheit = 40
	fmt.Println(celsius.ToFahrenheit())
	fmt.Println(fahrenheit.ToCelsius())
	fmt.Println(tempconv.IsBoiling(celsius))

}
