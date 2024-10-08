// Nilai celcius dalam fahrenheit, reamur, dan kelvin
package main

import "fmt"

func main() {
	var celcius, reamur, fahrenheit, kelvin float64
	fmt.Scan(&celcius)
	reamur = (4.0 / 5.0) * celcius
	fahrenheit = (9.0 / 5.0 * celcius) + 32
	kelvin = celcius + 273
	fmt.Println("Temperatur celcius :", celcius)
	fmt.Println("Derajat Reamur :", reamur)
	fmt.Println("Derajat Fahrenheit :", fahrenheit)
	fmt.Println("Derajat Kelvin :", kelvin)
}
