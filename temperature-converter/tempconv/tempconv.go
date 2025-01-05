package tempconv

type Celsius float64
type Fahrenheit float64

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius(f-32) * 5 / 9
}

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func IsBoiling(c Celsius) bool {
	return c >= 100
}
