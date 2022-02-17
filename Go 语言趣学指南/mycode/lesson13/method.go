package main

import "fmt"

// 声明类型
type celsius float64
type kelvin float64

// 开式度转换为摄氏度
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func main() {
	var k kelvin = 294.0
	c := k.celsius()
	fmt.Printf("%.2f°K is %.2f°C", k, c)

}
