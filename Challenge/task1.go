package Challenge

type Fahrenheit float32
type Celsius float32

func ToFahrenheit(t Celsius) Fahrenheit {
	var temp Fahrenheit
	temp = (Fahrenheit(t) * 9 / 5) + 32
	return temp
}
