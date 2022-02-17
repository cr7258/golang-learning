package main

import "fmt"

// 开式度转换为摄氏度
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

// 摄氏度转换为华氏度
func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32
}

// 开式度转换为华氏度
func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Printf("%.2f°K is %.2f°C\n", kelvin, celsius)
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("%.2f°C is %.2f°F\n", celsius, fahrenheit)
	fahrenheit = kelvinToFahrenheit(kelvin)
	fmt.Printf("%.2f°K is %.2f°F\n", kelvin, fahrenheit)
}
